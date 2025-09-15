# Cloudflare Zero Trust Access Policies for prototype-resort-apps.okiabrian.my.id

## Overview

Cloudflare Zero Trust Access provides identity-based security for your applications. This document details how to configure access policies for your resort apps prototype.

## Prerequisites

1. Cloudflare Zero Trust subscription
2. Domain configured in Cloudflare (prototype-resort-apps.okiabrian.my.id)
3. Identity provider (Google, GitHub, Microsoft, etc.)

## Step 1: Configure Identity Provider

### Google Identity Provider

1. Go to Zero Trust > Settings > Authentication
2. Click "Add new" under Login methods
3. Select "Google"
4. Configure:
   - Name: Google Auth
   - Client ID: [Your Google OAuth Client ID]
   - Client Secret: [Your Google OAuth Client Secret]
   - Redirect URL: Will be provided by Cloudflare

### GitHub Identity Provider

1. Go to Zero Trust > Settings > Authentication
2. Click "Add new" under Login methods
3. Select "GitHub"
4. Configure:
   - Name: GitHub Auth
   - Client ID: [Your GitHub OAuth Client ID]
   - Client Secret: [Your GitHub OAuth Client Secret]
   - Redirect URL: Will be provided by Cloudflare

## Step 2: Create Access Application

1. Go to Zero Trust > Access > Applications
2. Click "Add an application"
3. Select "Self-hosted"
4. Configure application details:
   - Application name: Resort Apps Prototype
   - Session duration: 24 hours
   - Application domain: prototype-resort-apps.okiabrian.my.id
   - Path: /*
   - Custom logo: Optional (you can upload your app logo)

## Step 3: Configure Application Settings

### Policy Configuration

1. In the application configuration, go to the "Policies" tab
2. Click "Add a policy"
3. Configure the main access policy:
   - Policy name: Allow Team Members
   - Action: Allow
   - Logic: Include
   - Selector: Emails
   - Value: [your-email@domain.com] (replace with actual email addresses)

### Additional Policies

You can create multiple policies for different access levels:

#### Policy for Admin Access
```
Policy name: Admin Access
Action: Allow
Logic: 
  - Include
    - Emails: [admin-email@domain.com]
  - AND
    - Groups: [admin-group]
```

#### Policy for Read-Only Access
```
Policy name: Read-Only Access
Action: Allow
Logic:
  - Include
    - Emails: [user1@domain.com, user2@domain.com]
```

## Step 4: Configure Service Auth for Backend

To secure communication between your frontend and backend:

1. Go to Zero Trust > Access > Service Auth
2. Click "Create a service token"
3. Name it "Resort Apps Backend"
4. Copy the Client ID and Client Secret
5. In your backend server, configure environment variables:
   ```
   CF_ACCESS_CLIENT_ID=[Client ID]
   CF_ACCESS_CLIENT_SECRET=[Client Secret]
   ```

6. Modify your Go server to validate service tokens:
   ```go
   // Add middleware to validate service tokens for API endpoints
   func validateServiceToken() gin.HandlerFunc {
       return func(c *gin.Context) {
           // Implementation to validate Cloudflare service tokens
           // This ensures only your frontend can access backend APIs
           c.Next()
       }
   }
   ```

## Step 5: Configure CORS and Headers

Your application already includes security headers, but with Cloudflare Access, you may need to adjust some settings:

1. In your application configuration, go to "Configure application"
2. Under "HTTP Headers", ensure these headers are set:
   ```
   X-Content-Type-Options: nosniff
   X-Frame-Options: DENY
   X-XSS-Protection: 1; mode=block
   Strict-Transport-Security: max-age=31536000; includeSubDomains
   Content-Security-Policy: default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'
   ```

## Step 6: Testing Access Policies

### Manual Testing

1. Visit https://prototype-resort-apps.okiabrian.my.id
2. You should be redirected to the Cloudflare Access login page
3. Try logging in with an authorized email
4. Try logging in with an unauthorized email (should be denied)

### Programmatic Testing

For API testing, you can obtain a service token:

1. Go to Zero Trust > Access > Service Auth
2. Create a test token
3. Use it in API requests:
   ```bash
   curl -H "CF-Access-Client-Id: [Client ID]" \
        -H "CF-Access-Client-Secret: [Client Secret]" \
        https://prototype-resort-apps.okiabrian.my.id/api/houses
   ```

## Step 7: Monitoring and Logs

### Access Logs

1. Go to Zero Trust > Logs > Access
2. Monitor:
   - Successful logins
   - Failed authentication attempts
   - Access denied events
   - Average session duration

### Setting Up Alerts

1. Go to Zero Trust > Logs > Settings
2. Configure alerts for:
   - High number of failed login attempts
   - Access from unusual locations
   - Multiple sessions from different IPs

## Advanced Configuration

### Device Posture Checks

Enhance security by requiring specific device configurations:

1. Go to Zero Trust > Devices > Device posture
2. Create posture checks:
   - OS version requirements
   - Disk encryption status
   - Firewall status
   - Antivirus status

### Geolocation Restrictions

Restrict access based on geographic location:

1. In your access policy, add:
   - Include: Country = [Your Country]
   - Exclude: Country = [High-risk Countries]

### Time-based Access

Control access based on time:

1. In your access policy, add time-based rules:
   - Allow access: Monday-Friday, 9AM-6PM
   - Deny access: Outside business hours

## Troubleshooting Common Issues

### Authentication Loop

If users are stuck in an authentication loop:

1. Check that the identity provider is correctly configured
2. Verify that the redirect URLs match exactly
3. Ensure the user's email is included in an access policy

### API Access Issues

If your frontend can't access backend APIs:

1. Verify service token configuration
2. Check that CORS headers allow the domain
3. Confirm that the backend is configured to accept requests from Cloudflare

### Mobile App Issues

For mobile applications:

1. Ensure the mobile user agent is not blocked
2. Verify that the mobile app can handle redirects properly
3. Test with different network conditions

This configuration provides robust identity-based security for your resort apps prototype while maintaining usability for authorized users.