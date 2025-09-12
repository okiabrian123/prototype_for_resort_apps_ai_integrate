#!/bin/bash

# Test server endpoints
echo "Testing health endpoint..."
curl -s http://localhost:8084/health
echo ""

echo "Testing root endpoint..."
curl -s http://localhost:8084/
echo ""

echo "Testing bookings endpoint..."
curl -s http://localhost:8084/api/bookings
echo ""
# Simple test script for the Resort App server

echo "Testing Resort App Server..."

# Test health endpoint
echo "Testing health endpoint..."
curl -s http://localhost:8080/health

echo ""

# Test main endpoint
echo "Testing main endpoint..."
curl -s http://localhost:8080/

echo ""

# Test bookings endpoint
echo "Testing bookings endpoint..."
curl -s http://localhost:8080/api/bookings

echo ""
echo "Server tests completed."