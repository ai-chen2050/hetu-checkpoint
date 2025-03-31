#!/bin/bash

# Default values
DEFAULT_VALIDATORS=1
CONFIG_DIR="docs/config"
BINARY_DIR="bin"

# Function to display usage
usage() {
    echo "Usage: $0 [-n number_of_validators] [-c config_dir] [-b binary_dir]"
    echo "  -n: Number of validators (default: ${DEFAULT_VALIDATORS})"
    echo "  -c: Config directory (default: ${CONFIG_DIR})"
    echo "  -b: Binary directory (default: ${BINARY_DIR})"
    exit 1
}

# Parse command line arguments
while getopts "n:c:b:h" opt; do
    case $opt in
    n) NUM_VALIDATORS=$OPTARG ;;
    c) CONFIG_DIR=$OPTARG ;;
    b) BINARY_DIR=$OPTARG ;;
    h) usage ;;
    ?) usage ;;
    esac
done

# Set number of validators to default if not specified
NUM_VALIDATORS=${NUM_VALIDATORS:-$DEFAULT_VALIDATORS}

# Validate number of validators
if ! [[ "$NUM_VALIDATORS" =~ ^[0-9]+$ ]] || [ "$NUM_VALIDATORS" -lt 1 ]; then
    echo "Error: Number of validators must be a positive integer"
    exit 1
fi

echo "Starting Hetu Checkpoint system..."
echo "Number of validators: $NUM_VALIDATORS"

# Create logs directory if it doesn't exist
mkdir -p logs

# Start dispatcher
echo "Starting dispatcher..."
"$BINARY_DIR"/dispatcher --config "$CONFIG_DIR"/dis_config.json >logs/dispatcher.log 2>&1 &
DISPATCHER_PID=$!
echo "Dispatcher started with PID: $DISPATCHER_PID"

# Wait for dispatcher to initialize
sleep 2

# Start validators
for ((i = 0; i < "$NUM_VALIDATORS"; i++)); do
    echo "Starting validator $i..."
    "$BINARY_DIR"/validator --config "$CONFIG_DIR"/val_config.json --id $i >logs/validator_$i.log 2>&1 &
    VALIDATOR_PIDS["$i"]=$!
    echo "Validator $i started with PID: ${VALIDATOR_PIDS[$i]}"
    sleep 1
done

# Create PID file
echo $DISPATCHER_PID >logs/dispatcher.pid
printf "%s\n" "${VALIDATOR_PIDS[@]}" >logs/validators.pid

echo "All processes started successfully!"
echo "Use './stop.sh' to stop all processes"

# Trap Ctrl+C and call cleanup
trap cleanup INT

cleanup() {
    echo "Stopping all processes..."
    kill $DISPATCHER_PID "${VALIDATOR_PIDS[@]}" 2>/dev/null
    exit 0
}

# Wait for any process to exit
wait
