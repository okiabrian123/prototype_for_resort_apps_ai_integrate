# Resort Chat Booking Application

A prototype for a mobile chat-based booking application for resorts, featuring natural language processing and a conversational UI.

## Project Overview

This application demonstrates a chat-based booking flow that guides users through three core questions:
1. When do you want to stay?
2. How many people?
3. Which house types are available?

The interface is designed to feel like a lightweight conversation rather than a long form, with a clean, minimal design featuring clear text, large buttons, and a friendly color palette.

## Current Status

This is a **60% complete prototype** created as a proof of concept for a freelance project bid. The UI components are complete and functional, and significant backend functionality has been implemented, including AI chatbot integration and booking data storage.

Potential clients can experience both the visual design and core functionality directly in their browser. The backend services require further configuration for full production functionality.

## Key Features Implemented

### UI Components (Complete)
- [x] Splash screen with fade animation
- [x] Home screen with resort promotions
- [x] Booking chat interface design
- [x] Order history screen
- [x] Notification system UI

### Backend Components (Mostly Complete)
- [x] Basic API structure with Go/Gin
- [x] Database setup with SQLite
- [x] Data models for houses and bookings
- [x] AI chatbot integration with OpenAI API
- [x] Function calling for house listings
- [x] Booking confirmation and data storage integration
- [x] Booking data processing and validation
- [ ] Complete booking management API
- [ ] Property management data integration

## Technology Stack

### Frontend
- React 19
- Tailwind CSS 4
- Vite 7
- React Router 7

### Backend
- Go 1.21
- Gin Framework
- SQLite database
- OpenAI API integration

## Project Structure

```
.
├── .env                     # Environment configuration (moved from server/.env)
├── resort-apps/             # Frontend React application
│   ├── src/
│   │   ├── components/
│   │   │   ├── SplashScreen.jsx
│   │   │   ├── HomeScreen.jsx
│   │   │   ├── BookingChatScreen.jsx
│   │   │   ├── OrderHistoryScreen.jsx
│   │   │   └── NotificationList.jsx
│   │   ├── App.jsx
│   │   └── main.jsx
│   ├── index.html
│   ├── package.json
│   ├── vite.config.js
│   └── tailwind.config.js
│
└── server/                  # Backend Go application
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
- Node.js (v14 or higher) for frontend development
- Go (v1.21 or higher) for backend services
- OpenAI API key for chatbot functionality (optional for demo)

### Environment Configuration

All environment variables are now managed in a single `.env` file in the root directory:

```bash
# Create the environment file from example
cp .env.example .env
```

Key environment variables:
- `DOMAIN_NAME`: The domain name for Cloudflare tunnel access
- `PORT`: Backend server port (default: 8084)
- `FRONTEND_PORT`: Frontend development server port (default: 5173)
- `ALLOWED_ORIGINS`: CORS allowed origins for security

### Frontend Installation

1. Navigate to the frontend directory:
   ```bash
   cd resort-apps
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

The frontend will be available at http://localhost:5173

### Backend Installation

1. Navigate to the backend directory:
   ```bash
   cd server
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Start the server:
   ```bash
   go run main.go
   ```

The backend will be available at http://localhost:8080 (default port)

### Running Both Services

To run both frontend and backend simultaneously, you can use the provided script:

#### On macOS/Linux:
```bash
./start-dev.sh
```

#### On Windows:
```bash
start-dev.bat
```

## Demo Access

Since this is a prototype for a mobile application, you can experience the UI design and user flow directly in your web browser without needing to download or install anything. The UI is designed to mimic a mobile app experience.

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

## UI Components

### Splash Screen
- Animated entry point with resort branding
- Automatic transition to home screen

### Home Screen
- Welcome message with resort branding
- Quick access to booking chat
- Order history navigation
- Notification system

### Booking Chat Screen
- Conversational interface design
- Message bubbles for user and bot
- Input area with send button
- Loading indicators
- House selection interface

### Order History Screen
- List of past bookings
- Status indicators for each booking
- Back navigation to home screen

## Backend Functionality

### AI Chatbot Integration
- OpenAI API integration for natural language processing
- Structured conversation flow for booking process
- Function calling for dynamic data retrieval
- House listing based on guest count
- Booking data extraction and validation

### Data Management
- SQLite database for persistent storage
- House data management
- Booking creation and retrieval
- Customer information handling

### Booking Confirmation System
- Automatic booking data extraction from AI responses
- Data validation before storage
- Database insertion with auto-generated IDs
- Booking status tracking (pending, confirmed, paid, cancelled)

### API Endpoints
- Resort/house management
- Booking management
- Chatbot AI integration
- Database operations

## Current Limitations

### AI Chatbot
- Requires OpenAI API key for full functionality
- System prompt configured for specific conversation flow
- Function calling implemented for house listings and booking data

### Data Management
- Sample data included for houses
- Database schema defined but limited sample data
- Booking validation in place but basic

## Future Development

### Backend Integration Completion
- Complete booking management API
- Property management data integration
- Enhanced error handling and validation
- Additional AI conversation flows

### Additional Features
- User authentication
- Payment processing integration
- Advanced search and filtering
- Admin dashboard for property management
- Multi-language support

## Contributing

This repository is primarily a prototype for a freelance project bid. As such, it is not actively seeking contributions. However, feedback and suggestions are welcome.

## Disclaimer

This is a 60% complete prototype created as part of a freelance project bid. It is not production-ready and is intended for demonstration purposes only. Several features are not yet implemented, and the code may require additional work to meet production standards.

## About the Developer

This prototype was created by an entry-level developer as a proof of concept for freelance work. The project demonstrates:
- Ability to create mobile-responsive UI interfaces
- Understanding of booking flow requirements for hospitality applications
- Proficiency in modern frontend development technologies (React, Tailwind CSS, Vite)
- Backend development skills (Go, Gin Framework, SQLite)
- AI integration capabilities (OpenAI API)

The 60% completion status shows commitment to proving capabilities rather than just submitting a proposal. Both the UI components and core backend functionality are complete and functional, allowing potential clients to experience the full booking flow.