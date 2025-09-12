#!/bin/bash

echo "Verifying that both servers are running..."

echo "Checking backend server (http://localhost:8084/health)..."
curl --max-time 5 -s http://localhost:8084/health || echo "Backend server not responding"

echo ""
echo "Checking frontend server (http://localhost:5173)..."
curl --max-time 5 -s -o /dev/null -w "HTTP Status: %{http_code}\n" http://localhost:5173

echo ""
echo "Servers verification complete."