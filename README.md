# Resort Chat Booking Application

A prototype for a mobile chat-based booking application for resorts, featuring natural language processing and a conversational UI.

## Project Overview

This application demonstrates a chat-based booking flow that guides users through three core questions:
1. When do you want to stay?
2. How many people?
3. Which house types are available?

The interface is designed to feel like a lightweight conversation rather than a long form, with a clean, minimal design featuring clear text, large buttons, and a friendly color palette.

## Current Status

This is a **60% complete prototype** created as a proof of concept for a freelance project bid. The core chat functionality and AI integration are working, allowing potential clients to experience the booking flow directly in their browser.

## Key Features Implemented

- [x] Splash screen with fade animation
- [x] Home screen with resort promotions
- [x] Booking chat interface with AI integration
- [x] Order history screen
- [x] Notification system
- [x] Real-time house availability based on guest count
- [x] Booking confirmation and data storage

## Technology Stack

### Frontend
- React 18
- Tailwind CSS
- Vite
- React Router

### Backend
- Go (Gin Framework)
- SQLite database
- OpenAI API integration

## Project Structure

```
resort-apps/
├── src/
│   ├── components/
│   │   ├── SplashScreen.jsx
│   │   ├── HomeScreen.jsx
│   │   ├── BookingChatScreen.jsx
│   │   ├── OrderHistoryScreen.jsx
│   │   └── NotificationList.jsx
│   ├── App.jsx
│   └── main.jsx
├── public/
├── index.html
├── package.json
└── tailwind.config.js

server/
├── data/
│   └── houses.json
├── database/
│   └── database.go
├── models/
│   ├── house.go
│   └── models.go
├── repository/
│   ├── bookings.go
│   └── resorts.go
├── tool_calling/
│   └── function_calling/
├── handlers.go
├── main.go
└── go.mod
```

## Quick Start

### Prerequisites
- Node.js (v14 or higher)
- Go (v1.21 or higher)
- npm or yarn

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd resort-apps
   ```

2. Install frontend dependencies:
   ```bash
   cd resort-apps
   npm install
   ```

3. Install backend dependencies:
   ```bash
   cd ../server
   go mod tidy
   ```

### Running the Application

#### Option 1: Using the convenience script (macOS/Linux)
```bash
cd ..
./start-dev.sh
```

#### Option 2: Manual start
1. Start the backend server:
   ```bash
   cd server
   go run *.go
   ```

2. Start the frontend development server (in a separate terminal):
   ```bash
   cd resort-apps
   npm run dev
   ```

The application will be available at:
- Frontend: http://localhost:5173
- Backend API: http://localhost:8084

## Demo Access

Since this is a prototype for a mobile application, you can experience the core functionality directly in your web browser without needing to download or install anything. The UI is designed to mimic a mobile app experience.

## Development

### Frontend Development
```bash
# Start development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

### Backend Development
```bash
# Run the server
go run main.go

# Run with specific port
PORT=8084 go run main.go
```

## API Endpoints

### Chat
- `POST /api/chat/message` - Send a message to the AI chatbot

### Houses
- `GET /api/houses` - Get all available houses
- `GET /api/houses/guests` - Get houses filtered by guest count
- `GET /api/houses/:id` - Get a specific house by ID
- `GET /api/houses/search/:query` - Search houses by name or location

### Bookings
- `GET /api/bookings` - Get all bookings
- `GET /api/bookings/:id` - Get a specific booking
- `POST /api/bookings` - Create a new booking
- `PUT /api/bookings/:id` - Update a booking
- `DELETE /api/bookings/:id` - Delete a booking

## Environment Variables

Create a `.env` file in the server directory based on `.env.example`:

```bash
# Server Configuration
PORT=8084
GIN_MODE=debug

# OpenAI Configuration (optional)
OPENAI_API_KEY=your_openai_api_key
OPENAI_BASE_URL=https://api.openai.com/v1
OPENAI_MODEL=gpt-3.5-turbo
```

## Database

The application uses SQLite for data persistence. The database file is automatically created at `server/data/resort.db` when the application starts.

## Contributing

This repository is primarily a prototype for a freelance project bid. As such, it is not actively seeking contributions. However, feedback and suggestions are welcome.

## Disclaimer

This is a prototype created as part of a freelance project bid. It is not production-ready and is intended for demonstration purposes only. Several features are not yet implemented, and the code may require additional work to meet production standards.

## About the Developer

This prototype was created by an entry-level developer as a proof of concept for freelance work. The project demonstrates:
- Ability to create chat-based UI interfaces
- Integration with AI APIs for natural language processing
- Full-stack development skills (React frontend with Go backend)
- Understanding of booking flow requirements for hospitality applications

The 60% completion status shows commitment to proving capabilities rather than just submitting a proposal.