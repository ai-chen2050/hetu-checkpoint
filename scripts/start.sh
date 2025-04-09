#!/bin/bash

# Default values
DEFAULT_VALIDATORS=1
CONFIG_DIR="docs/config"
BINARY_DIR="build"
KEYS_DIR="keys"
USE_DOCKER=false
REGISTER_AGAIN=false
DOCKER_COMPOSE_DISPATCHER="docker-compose-dispatcher.yml"
DOCKER_COMPOSE_VALIDATOR="docker-compose-validator.yml"

# Function to display usage
usage() {
    echo "Usage: $0 [-n number_of_validators] [-c config_dir] [-b binary_dir] [-k keys_dir] [-d]"
    echo "  -n: Number of validators (default: ${DEFAULT_VALIDATORS})"
    echo "  -c: Config directory (default: ${CONFIG_DIR})"
    echo "  -b: Binary directory (default: ${BINARY_DIR})"
    echo "  -k: Keys directory (default: ${KEYS_DIR})"
    echo "  -d: Use Docker Compose (default: false)"
    echo "  -r: Register and stake validator again (default: false)"
    echo "  -h: Display this help message"
    exit 1
}

# Parse command line arguments
while getopts "n:c:b:k:drh" opt; do
    case $opt in
    n) NUM_VALIDATORS=$OPTARG ;;
    c) CONFIG_DIR=$OPTARG ;;
    b) BINARY_DIR=$OPTARG ;;
    k) KEYS_DIR=$OPTARG ;;
    d) USE_DOCKER=true ;;
    r) REGISTER_AGAIN=true ;;
    h) usage ;;
    *) usage ;;
    esac
done

# Set number of validators if not specified
if [ -z "$NUM_VALIDATORS" ]; then
    NUM_VALIDATORS=$DEFAULT_VALIDATORS
fi

# Create logs directory if it doesn't exist
mkdir -p logs

# Function to start services using Docker Compose
start_with_docker() {
    echo "Starting services using Docker Compose..."

    # Check if key files exist, generate if not
    if [ ! -f "${KEYS_DIR}/dispatcher.json" ]; then
        echo "Generating dispatcher key..."
        mkdir -p "${KEYS_DIR}"
        "${BINARY_DIR}/dispatcher" generate-key --output="${KEYS_DIR}/dispatcher.json"
    fi

    # Start dispatcher
    echo "Starting dispatcher..."
    docker-compose -f "${DOCKER_COMPOSE_DISPATCHER}" up -d

    # Start validators
    for ((i = 0; i < "${NUM_VALIDATORS}"; i++)); do
        # Check if validator key exists, generate if not
        if [ ! -f "${KEYS_DIR}/validator_${i}.json" ]; then
            echo "Generating key for validator ${i}..."
            "${BINARY_DIR}/validator" generate-key --output="${KEYS_DIR}/validator_${i}.json"

            # Register and stake validator
            echo "Registering and staking validator ${i}..."
            "${BINARY_DIR}/validator" register-and-stake --config "${CONFIG_DIR}/val_config.json" --keys "${KEYS_DIR}/validator_${i}.json" --amount 500
        fi

        if [ "$REGISTER_AGAIN" = true ]; then
            # Register and stake validator
            echo "Registering and staking validator ${i}..."
            "${BINARY_DIR}/validator" register-and-stake --config "${CONFIG_DIR}/val_config.json" --keys "${KEYS_DIR}/validator_${i}.json" --amount 500
        fi

        # Start validator with Docker Compose
        echo "Starting validator ${i}..."
        # Create a copy of the docker-compose file with a unique name for each validator
        cp "${DOCKER_COMPOSE_VALIDATOR}" "docker-compose-validator-${i}.yml"

        # Update the validator ID and port in the docker-compose file
        # sed -i "s/--id 0/--id ${i}/g" "docker-compose-validator-${i}.yml"
        sed -i "s/- \"8081:8081\"/- \"$((8081 + i)):8081\"/g" "docker-compose-validator-${i}.yml"
        sed -i "s/- \"9001:9000\"/- \"$((9001 + i)):9000\"/g" "docker-compose-validator-${i}.yml"
        sed -i "s/- \"5433:5432\"/- \"$((5433 + i)):5432\"/g" "docker-compose-validator-${i}.yml"
        sed -i "s/container_name: portainer-validator/container_name: portainer-validator-${i}/g" "docker-compose-validator-${i}.yml"
        sed -i "s/portainer-validator-data/portainer-validator-${i}-data/g" "docker-compose-validator-${i}.yml"

        # Start the validator
        docker-compose -f "docker-compose-validator-${i}.yml" up -d
    done

    echo "All services started successfully!"
    echo "Use './scripts/stop.sh -d' to stop all services"
}

# Function to start services natively
start_natively() {
    echo "Starting services natively..."

    # Check if key files exist, generate if not
    if [ ! -f "${KEYS_DIR}/dispatcher.json" ]; then
        echo "Generating dispatcher key..."
        mkdir -p "${KEYS_DIR}"
        "${BINARY_DIR}/dispatcher" generate-key --output="${KEYS_DIR}/dispatcher.json"
    fi

    # Start dispatcher
    echo "Starting dispatcher..."
    "${BINARY_DIR}/dispatcher" run --config "${CONFIG_DIR}/dis_config.json" --keys="${KEYS_DIR}/dispatcher.json" --log-level=info --enable-db >logs/dispatcher.log 2>&1 &
    DISPATCHER_PID=$!
    echo "Dispatcher started with PID: $DISPATCHER_PID"

    # Create PID file for dispatcher
    echo $DISPATCHER_PID >logs/dispatcher.pid

    # Wait a moment for dispatcher to initialize
    sleep 2

    # Start validators
    declare -a VALIDATOR_PIDS
    for ((i = 0; i < "${NUM_VALIDATORS}"; i++)); do
        echo "Starting validator $i..."

        # Check if validator key exists, generate if not
        if [ ! -f "${KEYS_DIR}/validator_${i}.json" ]; then
            echo "Generating key for validator ${i}..."
            "${BINARY_DIR}/validator" generate-key --output="${KEYS_DIR}/validator_${i}.json"

            # Register and stake validator
            echo "Registering and staking validator ${i}..."
            "${BINARY_DIR}/validator" register-and-stake --config "${CONFIG_DIR}/val_config.json" --keys "${KEYS_DIR}/validator_${i}.json" --amount 500
        fi
        
        # Register and stake validator
        if [ "$REGISTER_AGAIN" = true ]; then
            echo "Registering and staking validator ${i}..."
            "${BINARY_DIR}/validator" register-and-stake --config "${CONFIG_DIR}/val_config.json" --keys "${KEYS_DIR}/validator_${i}.json" --amount 500
        fi

        # Start validator
        "${BINARY_DIR}/validator" run --config "${CONFIG_DIR}/val_config.json" --keys="${KEYS_DIR}/validator_${i}.json" --enable-db --log-level=info >logs/validator_${i}.log 2>&1 &
        VALIDATOR_PIDS["${i}"]=$!
        echo "Validator ${i} started with PID: ${VALIDATOR_PIDS[${i}]}"
        sleep 1
    done

    # Save validator PIDs to file
    printf "%s\n" "${VALIDATOR_PIDS[@]}" >logs/validators.pid

    echo "All processes started successfully!"
    echo "Use './scripts/stop.sh' to stop all processes"

    # Trap Ctrl+C and call cleanup
    trap cleanup INT

    cleanup() {
        echo "Stopping all processes..."
        kill "$DISPATCHER_PID" "${VALIDATOR_PIDS[@]}" 2>/dev/null
        exit 0
    }

    # Wait for any process to exit
    wait
}

# Start services based on the selected method
if [ "$USE_DOCKER" = true ]; then
    start_with_docker
else
    start_natively
fi