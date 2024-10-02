#!/bin/bash

# Check if port 8080 is in use
PID=$(lsof -t -i:8080)

# If PID is not empty
if [ -n "$PID" ]; then
    echo "Port 8080 is in use by process $PID. Attempting to kill..."

    # Attempt to kill the process
    kill -9 $PID

    # Check if the process was successfully killed
    if [ $? -eq 0 ]; then
        echo "Process $PID was successfully killed."
    else
        echo "Failed to kill process $PID. Please check your permissions."
    fi
else
    echo "Port 8080 is not in use."
fi