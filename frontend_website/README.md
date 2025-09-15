# Resort Apps

A beautiful resort booking application built with React and Tailwind CSS.

## Features

- Splash screen with fade animation
- Home screen with resort promotions
- Booking chat interface
- Order history screen

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm run dev
   ```

## Environment Configuration

The frontend now loads the allowed domain from the root `.env` file. The `DOMAIN_NAME` variable in the `.env` file is used to configure which domains are allowed to access the application through Vite's `allowedHosts` configuration.

## Backend Server

This application uses a Go backend server for handling bookings and order history.

### Server Setup

1. Navigate to the server directory:
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

The server will start on port 8080 by default.

## Running Both Servers

To run both the frontend and backend servers simultaneously, you can use the provided script:

### On macOS/Linux:
```bash
./start-dev.sh
```

### On Windows:
```bash
start-dev.bat
```

## Cloudflare Tunnel Deployment

For accessing the application through the Cloudflare tunnel at `prototype-resort-apps.okiabrian.my.id`, the Vite configuration has been updated to allow this host.

To preview the production build through the tunnel:

```bash
npm run build
npm run preview
```

## Technologies Used

- React 19
- React Router 7
- Tailwind CSS 4
- Vite 7
- Go (Gin Framework for backend)

## Project Structure

```
src/
  components/
    SplashScreen.jsx
    HomeScreen.jsx
    BookingChatScreen.jsx
    OrderHistoryScreen.jsx
  App.jsx
  main.jsx
  index.css

server/
  main.go
  handlers.go
  go.mod
  go.sum
```

## Development

To start the development server:

```bash
npm run dev
```

To build for production:

```bash
npm run build
```

To preview the production build:

```bash
npm run preview
```