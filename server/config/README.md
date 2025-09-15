# Cloudflare Zero Trust Subdomain Setup

This document explains how to set up the Cloudflare Zero Trust subdomain `prototype-resort-apps.okiabrian.my.id` using the Cloudflare CLI tool `cloudflared`.

## Overview

Cloudflare Zero Trust provides secure, authenticated access to applications without relying on traditional network perimeters. This setup uses Cloudflare Tunnel to expose your local services to the internet securely.

## Prerequisites

1. A Cloudflare account with Zero Trust subscription
2. Domain `okiabrian.my.id` configured in Cloudflare
3. [cloudflared](file:///root/prototype_for_resort_apps_ai_integrate/server/tool_calling/function_calling/cloudflared) CLI tool installed
4. The applications (frontend and backend) running locally

## Installation

### Install cloudflared

#### Linux (Debian/Ubuntu)
```bash
# Add Cloudflare's package signing key
sudo mkdir -p --mode=0755 /usr/share/keyrings
curl -fsSL https://pkg.cloudflare.com/cloudflare-main.gpg | sudo tee /usr/share/keyrings/cloudflare-main.gpg >/dev/null

# Add Cloudflare's apt repo
echo "deb [signed-by=/usr/share/keyrings/cloudflare-main.gpg] https://pkg.cloudflare.com/cloudflared any main" | sudo tee /etc/apt/sources.list.d/cloudflared.list

# Update and install
sudo apt-get update && sudo apt-get install cloudflared
```

#### macOS
```bash
brew install cloudflared
```

#### Windows
```powershell
winget install --id Cloudflare.cloudflared
```

#### Verify Installation
```bash
cloudflared --version
```

## Setup Process

### 1. Authenticate cloudflared
```bash
cloudflared tunnel login
```
This will open a browser window for you to log in to your Cloudflare account.

### 2. Create a Tunnel
```bash
cloudflared tunnel create prototype-resort-apps
```
This creates a tunnel named "prototype-resort-apps" and generates a credentials file.

### 3. Create Configuration File
Create a configuration file at `~/.cloudflared/config.yml`:

```yaml
tunnel: prototype-resort-apps
credentials-file: /home/user/.cloudflared/<TUNNEL-UUID>.json

# Ingress rules for routing traffic
ingress:
  # Frontend application (Vite dev server)
  - hostname: prototype-resort-apps.okiabrian.my.id
    service: http://localhost:5173
  # Backend API
  - hostname: prototype-resort-apps.okiabrian.my.id
    path: /api/
    service: http://localhost:8084
  # Backend API (alternative path)
  - hostname: prototype-resort-apps.okiabrian.my.id
    path: /health
    service: http://localhost:8084
  # Catch-all rule
  - service: http_status:404
```

Replace `<TUNNEL-UUID>` with the actual UUID from the previous step.

### 4. Route DNS
```bash
# Route the subdomain to your tunnel
cloudflared tunnel route dns prototype-resort-apps prototype-resort-apps.okiabrian.my.id
```

### 5. Configure Zero Trust Application

1. Log in to the Cloudflare Zero Trust dashboard
2. Go to Access > Applications
3. Click "Add an application"
4. Select "Self-hosted"
5. Configure:
   - Application name: Resort Apps Prototype
   - Domain: prototype-resort-apps.okiabrian.my.id
   - Path: /*
   - Session duration: 24 hours
6. Configure policies:
   - Policy name: Allow Authenticated Users
   - Action: Allow
   - Include:
     - Emails: [your-email@domain.com]
   - Authentication:
     - Select your preferred identity provider

### 6. Run the Tunnel
```bash
# Run the tunnel with the configuration
cloudflared tunnel --config ~/.cloudflared/config.yml run prototype-resort-apps
```

Or run by tunnel name:
```bash
cloudflared tunnel run prototype-resort-apps
```

## Environment Configuration

The setup script now supports custom environment configuration through the `server/.env` file. This allows you to customize ports and other settings without modifying the script directly.

### Available Environment Variables

- `PORT`: Port for the backend API service (default: 8084)
- `FRONTEND_PORT`: Port for the frontend development server (default: 5173)
- `DOMAIN_NAME`: Domain name for the application (default: prototype-resort-apps.okiabrian.my.id)
- `TUNNEL_NAME`: Name of the Cloudflare tunnel (default: prototype-resort-apps)
- `TUNNEL_UUID`: UUID of the Cloudflare tunnel (default: 9f8a5138-0c8e-4524-b6b8-287184532d72)

### Backend (.env)
```env
PORT=8084
GIN_MODE=debug
DB_PATH=./data/resort.db

# OpenAI Configuration
OPENAI_API_KEY=apikey-6d59e219d0264426bfa2c68fff983efe
OPENAI_BASE_URL=https://api.atlascloud.ai/v1
OPENAI_MODEL=zai-org/GLM-4.5-Air

# Security Configuration (Zero Trust) - Production settings
ALLOWED_ORIGINS=https://prototype-resort-apps.okiabrian.my.id
CORS_ENABLED=true

# Cloudflare Tunnel Configuration - Used by setup-cloudflare-tunnel.sh script
FRONTEND_PORT=5173
DOMAIN_NAME=prototype-resort-apps.okiabrian.my.id
TUNNEL_NAME=prototype-resort-apps
TUNNEL_UUID=9f8a5138-0c8e-4524-b6b8-287184532d72
```

### Frontend (Vite)
The frontend should automatically work with the tunnel setup.

## Automated Setup Script

For easier setup, you can use the automated script:

```bash
# Setup the tunnel (run once)
./setup-cloudflare-tunnel.sh setup

# Run the tunnel with automatic restart
./setup-cloudflare-tunnel.sh run

# Stop the tunnel
./setup-cloudflare-tunnel.sh stop

# Check tunnel status
./setup-cloudflare-tunnel.sh status
```

The script will automatically load configuration from `server/.env` if it exists.

## Useful Commands

### List Tunnels
```bash
cloudflared tunnel list
```

### Get Tunnel Information
```bash
cloudflared tunnel info prototype-resort-apps
```

### Stop Tunnel
Press `Ctrl+C` in the terminal where the tunnel is running.

### Delete Tunnel
```bash
cloudflared tunnel delete prototype-resort-apps
```

### Update cloudflared
```bash
cloudflared update
```

## Testing

1. Start your frontend and backend applications
2. Run the Cloudflare tunnel
3. Visit https://prototype-resort-apps.okiabrian.my.id
4. You should be redirected to the Cloudflare Access authentication page
5. After authentication, you should see your application

## Troubleshooting

### Common Issues

1. **DNS not resolving**: Ensure you've run the `cloudflared tunnel route dns` command
2. **Authentication loops**: Check your Zero Trust application configuration
3. **404 errors**: Verify your ingress rules in the config.yml file
4. **Connection refused**: Ensure your applications are running on the correct ports

### Logs
View tunnel logs:
```bash
cloudflared tunnel --loglevel debug run prototype-resort-apps
```

### Cleanup
If you need to clean up connections:
```bash
cloudflared tunnel cleanup prototype-resort-apps
```

## Security Considerations

1. Always use HTTPS (automatically provided by Cloudflare)
2. Configure strong authentication policies in Zero Trust
3. Regularly review and update access policies
4. Monitor logs for suspicious activity

This setup provides a secure way to expose your development applications to the internet using Cloudflare Zero Trust and Tunnel, with authentication required for access.