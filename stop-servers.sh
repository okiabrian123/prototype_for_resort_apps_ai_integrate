#!/bin/bash

# Script to stop the resort app servers

echo "Stopping resort app servers..."

# Kill processes by name
pkill -f "go run" >/dev/null 2>&1
pkill -f "npm run dev" >/dev/null 2>&1
pkill -f "cloudflared tunnel" >/dev/null 2>&1

# Also try to kill by PID if recorded
if [ -f /tmp/resort-app-pids ]; then
    source /tmp/resort-app-pids
    kill $BACKEND_PID 2>/dev/null
    kill $FRONTEND_PID 2>/dev/null
    if [ ! -z "$CLOUDFLARE_PID" ]; then
        kill $CLOUDFLARE_PID 2>/dev/null
    fi
    rm /tmp/resort-app-pids
fi

echo "Resort app servers stopped."