#!/bin/bash

echo "Stopping Hetu Checkpoint system..."

# Stop validators
if [ -f logs/validators.pid ]; then
    while read pid; do
        if kill -0 $pid 2>/dev/null; then
            echo "Stopping validator with PID: $pid"
            kill $pid
        fi
    done <logs/validators.pid
    rm logs/validators.pid
fi

# Stop dispatcher
if [ -f logs/dispatcher.pid ]; then
    DISPATCHER_PID=$(cat logs/dispatcher.pid)
    if kill -0 $DISPATCHER_PID 2>/dev/null; then
        echo "Stopping dispatcher with PID: $DISPATCHER_PID"
        kill $DISPATCHER_PID
    fi
    rm logs/dispatcher.pid
fi

echo "All processes stopped"
