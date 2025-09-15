# Cloudflare Zero Trust Setup for prototype-resort-apps.okiabrian.my.id

This document provides step-by-step instructions for setting up Cloudflare Zero Trust for your resort apps prototype domain.

## Overview

Your application consists of:
- Frontend: React/Vite application (port 5173)
- Backend: Go/Gin API server (port 8084)

We'll configure Cloudflare Zero Trust to secure both services with a zero-trust security model.

## Prerequisites

1. Cloudflare account with Zero Trust plan
2. Domain ownership of okiabrian.my.id
3. Access to DNS management for okiabrian.my.id
4. Server with public IP address or reverse proxy capability

## Step 1: Add Site to Cloudflare

1. Log in to your Cloudflare dashboard
2. Click "Add a site"
3. Enter your domain: `okiabrian.my.id`
4. Select the appropriate plan (Free plan works for basic setup)
5. Follow the DNS validation process

## Step 2: Configure DNS Records

Add the following DNS records in your Cloudflare DNS settings:

| Type | Name | Content | Proxy Status | TTL |
|------|------|---------|--------------|-----|
| A | prototype-resort-apps | [YOUR_SERVER_IP] | Proxied | Auto |
| CNAME | www.prototype-resort-apps | prototype-resort-apps.okiabrian.my.id | Proxied | Auto |

Note: Replace [YOUR_SERVER_IP] with your actual server's public IP address.

## Step 3: Configure SSL/TLS Encryption

1. In Cloudflare dashboard, go to SSL/TLS > Overview
2. Set SSL/TLS encryption mode to "Full (strict)" for maximum security
3. Go to SSL/TLS > Edge Certificates
4. Enable these settings:
   - Always Use HTTPS: On
   - Minimum TLS Version: TLS 1.2
   - Opportunistic Encryption: On
   - TLS 1.3: On
   - Automatic HTTPS Rewrites: On

## Step 4: Set up Cloudflare Zero Trust Access

1. Navigate to Zero Trust in your Cloudflare dashboard
2. Go to Access > Applications
3. Click "Add an application"
4. Select "Self-hosted"
5. Configure the application:
   - Application name: Resort Apps Prototype
   - Domain: prototype-resort-apps.okiabrian.my.id
   - Path: /*
   - Session duration: 24 hours

6. Configure policies:
   - Policy name: Allow Authenticated Users
   - Action: Allow
   - Include:
     - Emails: [your-email@domain.com] (replace with your email)
   - Authentication:
     - Select your preferred identity provider (Google, GitHub, etc.)

## Step 5: Configure Application Settings

### Backend API Configuration

1. In Access > Applications, click on your application
2. Go to "Configure application"
3. In the "Application URLs" section, add both services:
   - Main application: https://prototype-resort-apps.okiabrian.my.id
   - API service: https://prototype-resort-apps.okiabrian.my.id/api/

### CORS Configuration

Since your frontend makes API calls, you'll need to configure CORS properly:

1. In your server's .env file, ensure:
   ```
   ALLOWED_ORIGINS=https://prototype-resort-apps.okiabrian.my.id
   ```

2. Your server already includes security headers in the main.go file:
   ```go
   // Zero Trust Security: Set security headers
   c.Header("X-Content-Type-Options", "nosniff")
   c.Header("X-Frame-Options", "DENY")
   c.Header("X-XSS-Protection", "1; mode=block")
   c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
   c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'")
   ```

## Step 6: Firewall Rules Configuration

1. Go to Security > WAF > Tools in Cloudflare dashboard
2. Create firewall rules for additional protection:

### Rule 1: Block malicious user agents
- Field: User Agent
- Operator: matches regex
- Value: `(?:vuln\|nessus\|nikto\|nmap\|sqlmap\|dirbuster)`
- Action: Block

### Rule 2: Rate limiting
- Create a rate limiting rule under Security > Rate Limiting
- Match request URI: `*.okiabrian.my.id/*`
- Threshold: 1000 requests per 10 minutes
- Action: JS Challenge

## Step 7: Additional Security Headers

Add these additional security headers in Rules > Transform Rules:

1. Create a new transform rule
2. Add these response headers:
   ```
   Permissions-Policy: geolocation=(), microphone=(), camera=()
   Referrer-Policy: strict-origin-when-cross-origin
   ```

## Step 8: Deploy Your Application

There are two approaches to deploy your application with Cloudflare Zero Trust:

### Option 1: Traditional Deployment (Direct Server)
1. Build your frontend for production:
   ```bash
   cd resort-apps
   npm run build
   ```

2. Ensure your backend is configured for production:
   ```bash
   # In server/.env
   PORT=8084
   GIN_MODE=release
   ALLOWED_ORIGINS=https://prototype-resort-apps.okiabrian.my.id
   ```

3. Deploy both services to your server with direct IP access

### Option 2: Cloudflare Tunnel (Recommended for Development)
For development and testing, you can use Cloudflare Tunnel to securely expose your local services:

1. Use the automated setup script:
   ```bash
   ./setup-cloudflare-tunnel.sh setup
   ```

2. Run the tunnel with automatic restart capability:
   ```bash
   ./setup-cloudflare-tunnel.sh run
   ```

3. The tunnel will securely route traffic from `prototype-resort-apps.okiabrian.my.id` to your local services

## Environment Configuration

The Cloudflare tunnel script now supports custom environment configuration through the `server/.env` file. The following variables can be customized:

- `PORT`: Port for the backend API service (default: 8084)
- `FRONTEND_PORT`: Port for the frontend development server (default: 5173)
- `DOMAIN_NAME`: Domain name for the application (default: prototype-resort-apps.okiabrian.my.id)
- `TUNNEL_NAME`: Name of the Cloudflare tunnel (default: prototype-resort-apps)
- `TUNNEL_UUID`: UUID of the Cloudflare tunnel (default: 9f8a5138-0c8e-4524-b6b8-287184532d72)

To customize these values, modify the `server/.env` file before running the tunnel setup.

## Testing

After completing all steps:

1. Visit https://prototype-resort-apps.okiabrian.my.id
2. You should be redirected to the Cloudflare Access authentication page
3. After authentication, you should see your application
4. Test API calls to ensure they work correctly

## Troubleshooting

### Common Issues

1. **CORS errors**: Ensure ALLOWED_ORIGINS in your backend matches your domain
2. **Authentication loops**: Check that your identity provider is correctly configured
3. **Mixed content warnings**: Ensure all resources are loaded over HTTPS

### Logs and Monitoring

1. Check Cloudflare Analytics for traffic patterns
2. Monitor Access logs for authentication attempts
3. Review security events for any blocked requests

## Maintenance

1. Regularly review and update access policies
2. Monitor security logs for suspicious activity
3. Keep your origin server updated with security patches
4. Review and update firewall rules periodically

This setup provides a robust zero-trust security model for your prototype application while maintaining usability for authorized users.