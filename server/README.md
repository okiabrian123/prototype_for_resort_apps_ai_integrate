# Resort App Server(60% functionality)

This is the backend server for the Resort App, built with Go and Gin framework.

## Features

- RESTful API for booking management
- Resort listings management
- Health check endpoint
- Environment-based configuration
- SQLite database for data persistence
- AI Chatbot integration with OpenAI

## Prerequisites

- Go 1.21 or higher
- Gin framework
- godotenv for environment variables
- SQLite3
- OpenAI API key (for chatbot functionality)

## Installation

1. Navigate to the server directory:
   ```bash
   cd server
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Database Structure

### Bookings Table
| Column Name   | Type         | Description                              |
|---------------|--------------|------------------------------------------|
| id            | INTEGER      | Primary key (auto-increment)             |
| user_id       | INTEGER      | User identifier                          |
| resort_name   | TEXT         | Name of the resort                       |
| check_in      | DATE         | Check-in date                            |
| check_out     | DATE         | Check-out date                           |
| guests        | INTEGER      | Number of guests                         |
| total_price   | REAL         | Total price of booking                   |
| status        | TEXT         | Booking status (pending, confirmed, paid, cancelled) |
| payment_date  | DATE         | Payment date (optional)                  |
| created_at    | TIMESTAMP    | Creation timestamp                       |

## Configuration

Create a `.env` file based on the provided example:
```bash
cp .env.example .env
```

Update the values in `.env` according to your environment.

For the AI chatbot feature, you'll need to add your OpenAI API key:
```bash
OPENAI_API_KEY=your_openai_api_key_here
```

If you're using a custom OpenAI-compatible API endpoint, you can also set:
```bash
OPENAI_BASE_URL=https://api.openai.com/v1
```

If you want to specify a particular model (especially important for custom endpoints):
```bash
OPENAI_MODEL=openai/gpt-3.5-turbo
```

## Running the Server

```bash
go run main.go
```

The server will start on port 8084 by default (http://localhost:8084).

## API Endpoints

### Health Check
- `GET /health` - Server health status

### Resorts
- `GET /api/resorts` - Get all resorts
- `GET /api/resorts/:id` - Get a specific resort
- `GET /api/resorts/search/:query` - Search resorts by name or location

### Bookings
- `GET /api/bookings` - Get all bookings
- `GET /api/bookings/:id` - Get a specific booking
- `GET /api/bookings/status/:status` - Get bookings by status
- `GET /api/bookings/user/:user_id` - Get bookings by user ID
- `POST /api/bookings` - Create a new booking
- `PUT /api/bookings/:id` - Update a booking
- `DELETE /api/bookings/:id` - Delete a booking

### Chatbot
- `POST /api/chat/message` - Send a message to the AI chatbot

## Booking Status Values

- `pending` - Booking created but not confirmed
- `confirmed` - Booking confirmed but not paid
- `paid` - Booking confirmed and paid
- `cancelled` - Booking cancelled

## AI Chatbot

The resort app includes an AI chatbot powered by OpenAI's GPT models. The chatbot can assist users with:

- Booking inquiries
- Resort information
- Availability questions
- General assistance

For detailed information about the chatbot implementation, see [README_CHATBOT.md](README_CHATBOT.md).