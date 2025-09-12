@echo off
REM Script to start both frontend and backend servers on Windows

echo Starting Resort App development environment...

REM Start the Go backend server in the background
echo Starting backend server...
cd /D "%~dp0server"
start "Backend Server" go run *.go

REM Give the backend server a moment to start
timeout /t 3 /nobreak >nul

echo Backend server started

REM Start the React frontend
echo Starting frontend server...
cd /D "%~dp0resort-apps"
start "Frontend Server" npm run dev

echo Frontend server started
echo Servers are running. Close the terminal windows to stop.

REM Keep the script running
pause