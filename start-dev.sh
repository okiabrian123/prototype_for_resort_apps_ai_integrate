#!/bin/bash

# Script to start both frontend and backend servers
# Kills any existing processes first

# Default values
CLOUDFLARE_MODE=false

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --cloudflare)
            CLOUDFLARE_MODE=true
            shift
            ;;
        -*)
            echo "Unknown option $1"
            exit 1
            ;;
        *)
            # Positional arguments (ports)
            if [ -z "$POSITIONAL_ARGS" ]; then
                POSITIONAL_ARGS="$1"
            else
                POSITIONAL_ARGS="$POSITIONAL_ARGS $1"
            fi
            shift
            ;;
    esac
done

# Set positional arguments if any
if [ -n "$POSITIONAL_ARGS" ]; then
    set -- $POSITIONAL_ARGS
fi

# Load environment variables from .env file only
if [ -f ".env" ]; then
    echo "Loading environment variables from .env file"
    export $(grep -v '^#' .env | xargs)
else
    echo "Error: .env file not found. Please create .env file with required configuration."
    exit 1
fi

# Check if port arguments were provided
if [ $# -eq 0 ]; then
    BACKEND_PORT=$PORT
    FRONTEND_PORT=$FRONTEND_PORT
elif [ $# -eq 1 ]; then
    BACKEND_PORT=$1
    FRONTEND_PORT=$FRONTEND_PORT
else
    BACKEND_PORT=$1
    FRONTEND_PORT=$2
fi

# Verify that required environment variables are set
if [ -z "$PORT" ] || [ -z "$FRONTEND_PORT" ]; then
    echo "Error: PORT and FRONTEND_PORT must be defined in .env file"
    exit 1
fi

echo "Stopping any existing server processes..."
pkill -f "go run" >/dev/null 2>&1
pkill -f "npm run dev" >/dev/null 2>&1

# Additional check for processes using our ports
echo "Checking for processes on ports $BACKEND_PORT and $FRONTEND_PORT..."
lsof -i :$BACKEND_PORT | grep LISTEN | awk '{print $2}' | xargs kill -9 >/dev/null 2>&1
lsof -i :$FRONTEND_PORT | grep LISTEN | awk '{print $2}' | xargs kill -9 >/dev/null 2>&1

sleep 3

echo "Starting backend server on port $BACKEND_PORT..."
cd server
PORT=$BACKEND_PORT go run *.go &
BACKEND_PID=$!
cd ..

echo "Starting frontend server on port $FRONTEND_PORT..."
cd frontend_website
npm run dev -- --port $FRONTEND_PORT &
FRONTEND_PID=$!
cd ..

# Run Cloudflare tunnel if requested
if [ "$CLOUDFLARE_MODE" = true ]; then
    echo "Starting Cloudflare tunnel..."
    ./setup-cloudflare-tunnel.sh run &
    CLOUDFLARE_PID=$!
    echo "Cloudflare tunnel PID: $CLOUDFLARE_PID"
fi

echo "Servers started!"
echo "Backend PID: $BACKEND_PID (http://localhost:$BACKEND_PORT)"
echo "Frontend PID: $FRONTEND_PID (http://localhost:$FRONTEND_PORT)"
if [ "$CLOUDFLARE_MODE" = true ]; then
    echo "Cloudflare tunnel PID: $CLOUDFLARE_PID (https://$DOMAIN_NAME)"
fi
echo ""
echo "Press Ctrl+C to stop all processes"

# Cleanup function to kill all processes on exit
cleanup() {
    echo "Stopping all processes..."
    kill $BACKEND_PID $FRONTEND_PID 2>/dev/null
    if [ "$CLOUDFLARE_MODE" = true ]; then
        kill $CLOUDFLARE_PID 2>/dev/null
    fi
    wait $BACKEND_PID $FRONTEND_PID 2>/dev/null
    if [ "$CLOUDFLARE_MODE" = true ]; then
        wait $CLOUDFLARE_PID 2>/dev/null
    fi
    exit 0
}

# Trap SIGINT (Ctrl+C) to clean up processes
trap cleanup SIGINT

# Wait for all processes
wait $BACKEND_PID $FRONTEND_PID
if [ "$CLOUDFLARE_MODE" = true ]; then
    wait $CLOUDFLARE_PID
fi