# AI Chatbot Setup Guide

This guide explains how to set up and use the AI chatbot feature that has been implemented in the resort booking application.

## Overview

The AI chatbot feature provides users with an intelligent assistant that can answer questions about resort bookings, provide information about available resorts, and guide users through the booking process.

## Prerequisites

Before setting up the chatbot, ensure you have:

1. An OpenAI API key (sign up at https://platform.openai.com/)
2. Go 1.21 or higher installed
3. Node.js and npm installed
4. All existing project dependencies installed

## Setup Instructions

### 1. Configure OpenAI API Key

Add your OpenAI API key to the server environment:

1. Open the `.env` file in the `server` directory
2. Add the following line:
   ```
   OPENAI_API_KEY=your_actual_openai_api_key_here
   ```
3. Replace `your_actual_openai_api_key_here` with your actual API key

### 2. (Optional) Configure Custom OpenAI-Compatible API Endpoint

If you're using a custom OpenAI-compatible API endpoint (such as OpenRouter, Together.ai, etc.), you can set the base URL:

1. In the same `.env` file, add:
   ```
   OPENAI_BASE_URL=https://api.openai.com/v1
   ```
2. Replace the URL with your custom endpoint

### 3. Install Dependencies

1. Navigate to the server directory:
   ```bash
   cd server
   ```

2. Run go mod tidy to install the OpenAI library:
   ```bash
   go mod tidy
   ```

### 4. Start the Application

1. Start the backend server:
   ```bash
   cd server
   go run *.go
   ```

2. In a separate terminal, start the frontend:
   ```bash
   cd resort-apps
   npm run dev
   ```

## Using the Chatbot

1. Open your browser and navigate to the application (typically http://localhost:5173 or similar)
2. Navigate to the Booking Chat screen
3. Type your message in the input field at the bottom
4. Press Enter or click the Send button
5. Wait for the AI assistant to respond
6. Continue the conversation as needed

## Features

The AI chatbot includes the following features:

- Real-time conversation with an AI assistant
- Context-aware responses based on conversation history
- Loading indicators while the AI is processing
- Error handling for API failures
- Professional resort booking assistant persona
- Support for custom OpenAI-compatible API endpoints

## Customization

You can customize the AI assistant's behavior by modifying the system message in the `chatWithAI` function in `server/handlers.go`:

```go
systemMessage := openai.ChatCompletionMessage{
    Role:    openai.ChatMessageRoleSystem,
    Content: "Your custom instructions here",
}
```

## Troubleshooting

### Common Issues

1. **API Key Error**: Ensure your OpenAI API key is correctly set in the `.env` file
2. **Network Issues**: Check that the frontend can reach the backend server
3. **Rate Limiting**: OpenAI has rate limits; implement appropriate error handling

### Testing the API

You can test the chat endpoint directly with curl:

```bash
curl -X POST http://localhost:8084/api/chat/message \
  -H "Content-Type: application/json" \
  -d '{
    "messages": [
      {
        "role": "user",
        "content": "Hello, I would like to book a resort for 2 people."
      }
    ]
  }'
```

## Files Modified

The following files were modified to implement the chatbot feature:

### Backend
- `server/go.mod` - Added OpenAI dependency
- `server/.env` - Added API key and base URL configuration
- `server/main.go` - Added chat API endpoint
- `server/handlers.go` - Added chat handler function with custom URL support

### Frontend
- `resort-apps/src/components/BookingChatScreen.jsx` - Updated component for AI integration

## Architecture

The chatbot follows a client-server architecture:

1. **Frontend**: React component that provides the chat interface
2. **Backend**: Go server that handles API requests and communicates with OpenAI
3. **OpenAI**: Processes natural language and generates responses

For more detailed technical information, see:
- [CHATBOT_IMPLEMENTATION_SUMMARY.md](CHATBOT_IMPLEMENTATION_SUMMARY.md)
- [server/README_CHATBOT.md](server/README_CHATBOT.md)
- [server/CHATBOT_ARCHITECTURE.md](server/CHATBOT_ARCHITECTURE.md)