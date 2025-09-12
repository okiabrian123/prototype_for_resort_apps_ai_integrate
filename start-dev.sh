#!/bin/bash

# Script to start both frontend and backend servers
# Kills any existing processes first

echo "Stopping any existing server processes..."
pkill -f "go run" >/dev/null 2>&1
pkill -f "npm run dev" >/dev/null 2>&1

# Additional check for processes using our ports
echo "Checking for processes on ports 8084 and 5173..."
lsof -i :8084 | grep LISTEN | awk '{print $2}' | xargs kill -9 >/dev/null 2>&1
lsof -i :5173 | grep LISTEN | awk '{print $2}' | xargs kill -9 >/dev/null 2>&1

sleep 3

echo "Starting backend server on port 8084..."
cd server
go run *.go &
BACKEND_PID=$!
cd ..

echo "Starting frontend server..."
cd resort-apps
npm run dev &
FRONTEND_PID=$!
cd ..

echo "Servers started!"
echo "Backend PID: $BACKEND_PID (http://localhost:8084)"
echo "Frontend PID: $FRONTEND_PID (http://localhost:5173)"
echo ""
echo "Press Ctrl+C to stop both servers"

# Wait for both processes
wait $BACKEND_PID $FRONTEND_PID