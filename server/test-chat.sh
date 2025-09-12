#!/bin/bash

# Test the chat endpoint
curl -X POST http://localhost:8084/api/chat/message \
  -H "Content-Type: application/json" \
  -d '{
    "messages": [
      {
        "role": "user",
        "content": "Hello, what can you help me with?"
      }
    ]
  }'