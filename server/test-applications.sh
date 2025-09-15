#!/bin/bash

# Test script to verify that the resort applications are running

echo "Testing Resort Applications"
echo "=========================="

# Test backend health endpoint
echo "1. Testing backend health endpoint..."
if curl -s http://localhost:8084/health | grep -q "healthy"; then
    echo "✓ Backend is running and healthy"
else
    echo "✗ Backend is not responding or not healthy"
fi

# Test backend API
echo "2. Testing backend API..."
if curl -s http://localhost:8084/api/houses | grep -q "houses"; then
    echo "✓ Backend API is accessible"
else
    echo "✗ Backend API is not accessible"
fi

# Test frontend (simple check)
echo "3. Testing frontend..."
if curl -s http://localhost:5173 | grep -q "html"; then
    echo "✓ Frontend is running"
else
    echo "✗ Frontend is not responding"
fi

echo ""
echo "Make sure both applications are running before starting the Cloudflare Tunnel:"
echo "- Frontend (Vite): npm run dev (in resort-apps directory)"
echo "- Backend (Go): go run main.go (in server directory)"