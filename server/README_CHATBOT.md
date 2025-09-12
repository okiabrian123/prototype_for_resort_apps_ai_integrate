# AI Chatbot Integration with OpenAI

This document explains how to set up and use the AI chatbot feature integrated with OpenAI in the resort booking application.

## Setup Instructions

### 1. Obtain an OpenAI API Key

1. Visit the OpenAI website at https://platform.openai.com/account/api-keys
2. If you don't have an account, click on "Sign Up" to create one
3. Once logged in, navigate to your API key management page
4. Click on "Create new secret key"
5. Copy the generated API key

### 2. Configure the Application

1. Update the `.env` file in the server directory with your OpenAI API key:
   ```
   OPENAI_API_KEY=your_actual_api_key_here
   ```
   
2. Optionally, if you're using a custom OpenAI-compatible API endpoint, set the base URL:
   ```
   OPENAI_BASE_URL=https://api.openai.com/v1
   ```
   
3. If you want to use a specific model (especially important for custom endpoints), set the model:
   ```
   OPENAI_MODEL=openai/gpt-3.5-turbo
   ```

3. The required dependency has already been added to `go.mod`:
   ```
   github.com/sashabaranov/go-openai v1.28.1
   ```

### 3. Run the Application

1. Start the backend server:
   ```bash
   cd server
   go run *.go
   ```

2. Start the frontend (in a separate terminal):
   ```bash
   cd resort-apps
   npm run dev
   ```

## API Endpoints

### Chat Endpoint

- **URL**: `/api/chat/message`
- **Method**: `POST`
- **Body**: 
  ```json
  {
    "messages": [
      {
        "role": "user|assistant",
        "content": "Message content"
      }
    ]
  }
  ```
- **Response**:
  ```json
  {
    "message": "AI response content"
  }
  ```

## Frontend Integration

The BookingChatScreen component has been updated to communicate with the backend AI service:

1. Messages are sent to the backend when the user submits them
2. The backend forwards the conversation to OpenAI
3. The AI response is displayed in the chat interface
4. Loading indicators show when the AI is processing
5. Error handling for API failures