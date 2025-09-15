#!/bin/bash

# Script to start both frontend and backend servers in background mode for crontab
# This script is designed to be run from crontab and will not block

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# Add Go binary path to ensure commands can be found
export PATH="/usr/local/go/bin:$PATH"

# Load environment variables from .env file
if [ -f ".env" ]; then
    echo "$(date): Loading environment variables from .env file" >> /var/log/resort-app-startup.log
    export $(grep -v '^#' .env | xargs)
else
    echo "$(date): Error: .env file not found" >> /var/log/resort-app-startup.log
    exit 1
fi

# Verify that required environment variables are set
if [ -z "$PORT" ] || [ -z "$FRONTEND_PORT" ]; then
    echo "$(date): Error: PORT and FRONTEND_PORT must be defined in .env file" >> /var/log/resort-app-startup.log
    exit 1
fi

# Set the ports
BACKEND_PORT=$PORT
FRONTEND_PORT=$FRONTEND_PORT

echo "$(date): Stopping any existing server processes..." >> /var/log/resort-app-startup.log
pkill -f "go run" >/dev/null 2>&1
pkill -f "npm run dev" >/dev/null 2>&1

# Additional check for processes using our ports
lsof -i :$BACKEND_PORT | grep LISTEN | awk '{print $2}' | xargs kill -9 >/dev/null 2>&1
lsof -i :$FRONTEND_PORT | grep LISTEN | awk '{print $2}' | xargs kill -9 >/dev/null 2>&1

sleep 3

echo "$(date): Starting backend server on port $BACKEND_PORT..." >> /var/log/resort-app-startup.log
cd server
PORT=$BACKEND_PORT nohup /usr/local/go/bin/go run *.go > /var/log/resort-backend.log 2>&1 &
BACKEND_PID=$!
cd ..

echo "$(date): Starting frontend server on port $FRONTEND_PORT..." >> /var/log/resort-app-startup.log
cd frontend_website
nohup npx npm run dev -- --port $FRONTEND_PORT > /var/log/resort-frontend.log 2>&1 &
FRONTEND_PID=$!
cd ..

# Run Cloudflare tunnel if requested
if [ "$1" = "--cloudflare" ]; then
    echo "$(date): Starting Cloudflare tunnel..." >> /var/log/resort-app-startup.log
    nohup ./setup-cloudflare-tunnel.sh run > /var/log/resort-cloudflare.log 2>&1 &
    CLOUDFLARE_PID=$!
    echo "$(date): Cloudflare tunnel PID: $CLOUDFLARE_PID" >> /var/log/resort-app-startup.log
fi

echo "$(date): Servers started!" >> /var/log/resort-app-startup.log
echo "$(date): Backend PID: $BACKEND_PID (http://localhost:$BACKEND_PORT)" >> /var/log/resort-app-startup.log
echo "$(date): Frontend PID: $FRONTEND_PID (http://localhost:$FRONTEND_PORT)" >> /var/log/resort-app-startup.log
if [ "$1" = "--cloudflare" ]; then
    echo "$(date): Cloudflare tunnel PID: $CLOUDFLARE_PID" >> /var/log/resort-app-startup.log
fi

# Write PIDs to file for later management
echo "BACKEND_PID=$BACKEND_PID" > /tmp/resort-app-pids
echo "FRONTEND_PID=$FRONTEND_PID" >> /tmp/resort-app-pids
if [ "$1" = "--cloudflare" ]; then
    echo "CLOUDFLARE_PID=$CLOUDFLARE_PID" >> /tmp/resort-app-pids
fi