#!/bin/bash

## Function to display usage information
usage() {
  cat << EOF
Usage: $0 [OPTIONS] [ARGUMENTS]

A demo script showing advanced getopt features.

Options:
  -i, --input FILE       Input file to process (required)
  -o, --output FILE      Output file (default: output.txt)
  -m, --mode MODE        Processing mode: normal|verbose (default: normal)
  -l, --log FILE         Log file (default: none)
  -v, --verbose          Enable verbose output
  -h, --help             Display this help message

Examples:
  $0 -i input.txt -o output.txt
  $0 --input=data.csv --mode=verbose
  $0 -i input.txt -v -l log.txt
EOF
  exit 1
}

## Function to log messages
log_message() {
  local message="$1"
  local timestamp=$(date "+%Y-%m-%d %H:%M:%S")

  echo "[$timestamp] $message"

  if [ -n "$LOG_FILE" ]; then
    echo "[$timestamp] $message" >> "$LOG_FILE"
  fi
}

## Function to process a file
process_file() {
  local input="$1"
  local output="$2"
  local mode="$3"

  if [ ! -f "$input" ]; then
    log_message "Error: Input file '$input' does not exist."
    return 1
  fi

  log_message "Processing file: $input"
  log_message "Output file: $output"
  log_message "Mode: $mode"

  ## Perform different operations based on mode
  if [ "$mode" = "verbose" ]; then
    log_message "File details:"
    log_message "  - Size: $(wc -c < "$input") bytes"
    log_message "  - Lines: $(wc -l < "$input") lines"
    log_message "  - Words: $(wc -w < "$input") words"
  fi

  ## Simulate processing
  log_message "Reading input file..."
  cat "$input" > "$output"
  log_message "Processing complete."
  log_message "Output written to: $output"

  return 0
}

## Parse command-line options
OPTS=$(getopt -o i:o:m:l:vh --long input:,output:,mode:,log:,verbose,help -n 'advanced_getopt.sh' -- "$@")

if [ $? -ne 0 ]; then
  echo "Failed to parse options" >&2
  usage
fi

## Reset the positional parameters to the parsed options
eval set -- "$OPTS"

## Initialize variables with default values
INPUT_FILE=""
OUTPUT_FILE="output.txt"
MODE="normal"
LOG_FILE=""
VERBOSE=false

## Process the options
while true; do
  case "$1" in
    -i | --input)
      INPUT_FILE="$2"
      shift 2
      ;;
    -o | --output)
      OUTPUT_FILE="$2"
      shift 2
      ;;
    -m | --mode)
      if [ "$2" = "normal" ] || [ "$2" = "verbose" ]; then
        MODE="$2"
      else
        echo "Error: Invalid mode '$2'. Must be 'normal' or 'verbose'." >&2
        usage
      fi
      shift 2
      ;;
    -l | --log)
      LOG_FILE="$2"
      shift 2
      ;;
    -v | --verbose)
      VERBOSE=true
      shift
      ;;
    -h | --help)
      usage
      ;;
    --)
      shift
      break
      ;;
    *)
      echo "Internal error!"
      exit 1
      ;;
  esac
done

## Check if required options are provided
if [ -z "$INPUT_FILE" ]; then
  echo "Error: Input file must be specified with -i or --input option." >&2
  usage
fi

## Enable verbose mode if specified
if [ "$VERBOSE" = true ] && [ "$MODE" != "verbose" ]; then
  MODE="verbose"
fi

## Process the file
process_file "$INPUT_FILE" "$OUTPUT_FILE" "$MODE"
EXIT_CODE=$?

## Additional arguments are available as $1, $2, etc.
if [ $# -gt 0 ]; then
  log_message "Additional arguments provided: $@"
fi

exit $EXIT_CODE
