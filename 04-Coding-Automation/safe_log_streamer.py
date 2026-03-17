"""
Google SRE Coding Pattern: Safe Log Streamer

THE SCENARIO:
You are asked to parse a 500GB NDJSON (Newline Delimited JSON) access log, 
count the HTTP status codes, and return a summary.

THE TRAP (L3/Junior):
Using `open(file).read()` or `file.readlines()`. This loads 500GB into RAM, 
triggers the Linux OOM Killer, and takes down the production server.

THE SRE SIGNAL (L4/L5):
1. O(1) Memory: Use generators/iterators to stream the file line-by-line.
2. Defensive Parsing: Expect corrupted data. Never crash on a bad log line.
3. Observability: Count the malformed lines. If 100% of lines are bad, alert the operator.
4. Unix Philosophy: Support `stdin` so the script can be piped (e.g., `zcat logs.gz | ./script.py`).
"""

import sys
import json
from collections import Counter
import logging

# Configure basic logging for the operator (Stderr so it doesn't pollute stdout data)
logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")

def stream_and_aggregate(file_stream, error_threshold_percent=10.0):
    """
    Streams a file object, aggregates status codes, and tracks malformed lines.
    Maintains an O(1) memory footprint for the file buffer.
    """
    status_counts = Counter()
    metrics = {
        "total_lines_processed": 0,
        "valid_lines": 0,
        "malformed_lines": 0
    }

    # SRE SIGNAL: Lazy evaluation. We process one line in memory at a time.
    for line_number, line in enumerate(file_stream, start=1):
        line = line.strip()
        if not line:
            continue  # Skip empty lines gracefully
        
        metrics["total_lines_processed"] += 1

        # SRE SIGNAL: Defensive Programming. 
        # Production logs are frequently corrupted, truncated, or missing fields.
        try:
            log_entry = json.loads(line)
            status_code = log_entry.get("status")
            
            if not status_code:
                # Log lacks the required schema, treat as malformed
                metrics["malformed_lines"] += 1
                continue
                
            status_counts[status_code] += 1
            metrics["valid_lines"] += 1

        except json.JSONDecodeError:
            # SRE SIGNAL: Never crash the pipeline on a bad log line.
            # Record the metric so operators have visibility into log quality.
            metrics["malformed_lines"] += 1

    # SRE SIGNAL: Sanity Check / Circuit Breaker
    # If the schema changed upstream, 100% of lines might be "malformed".
    # We must catch this and fail loudly, rather than silently reporting "0 status codes".
    if metrics["total_lines_processed"] > 0:
        error_rate = (metrics["malformed_lines"] / metrics["total_lines_processed"]) * 100
        if error_rate > error_threshold_percent:
            logging.error(f"Data quality alert: {error_rate:.2f}% of logs were malformed "
                          f"(Threshold is {error_threshold_percent}%). Upstream schema may have changed.")
            sys.exit(1) # SRE SIGNAL: Return non-zero exit code for CI/CD or Cron failure detection

    return status_counts, metrics


def main():
    # SRE SIGNAL: Support stdin for pipeline composition
    # Usage: cat access.jsonl | python3 safe_log_streamer.py
    
    # Check if data is being piped in
    if not sys.stdin.isatty():
        input_stream = sys.stdin
    else:
        # Fallback for demonstration if run without pipe
        logging.info("No stdin detected. Please pipe data: `cat logs.jsonl | python3 safe_log_streamer.py`")
        sys.exit(0)

    status_counts, metrics = stream_and_aggregate(input_stream)

    # SRE SIGNAL: Output structured data (JSON) so downstream tools can parse it
    final_output = {
        "aggregation": dict(status_counts),
        "observability_metrics": metrics
    }
    
    # Print to stdout (the actual output of the tool)
    print(json.dumps(final_output, indent=2))

if __name__ == "__main__":
    main()
