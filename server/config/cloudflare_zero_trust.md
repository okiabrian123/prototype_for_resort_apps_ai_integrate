# Cloudflare Zero Trust Subdomain Setup for prototype-resort-apps.okiabrian.my.id

This document provides specific instructions for setting up the Cloudflare Zero Trust subdomain for the resort application.

## Subdomain Configuration

The subdomain `prototype-resort-apps.okiabrian.my.id` has been configured with the following settings:

### DNS Configuration
- **Type**: A Record
- **Name**: prototype-resort-apps
- **Content**: [YOUR_SERVER_IP] (Replace with actual server IP)
- **Proxy Status**: Proxied
- **TTL**: Auto

Optional WWW Record:
- **Type**: CNAME
- **Name**: www.prototype-resort-apps
- **Content**: prototype-resort-apps.okiabrian.my.id
- **Proxy Status**: Proxied
- **TTL**: Auto

### SSL/TLS Configuration
- **Encryption Mode**: Full (strict)
- **Always Use HTTPS**: On
- **Minimum TLS Version**: TLS 1.2
- **TLS 1.3**: On
- **Automatic HTTPS Rewrites**: On

### Zero Trust Access Application

Application Settings:
- **Name**: Resort Apps Prototype
- **Domain**: prototype-resort-apps.okiabrian.my.id
- **Path**: /*
- **Session Duration**: 24 hours

Access Policy:
- **Policy Name**: Allow Authenticated Users
- **Action**: Allow
- **Include**:
  - Emails: [your-email@domain.com] (Replace with authorized email addresses)
- **Authentication**:
  - Identity Provider: Google/GitHub/Microsoft (Configure as needed)

### Backend API Configuration

API Service URLs:
- **Main Application**: https://prototype-resort-apps.okiabrian.my.id
- **API Service**: https://prototype-resort-apps.okiabrian.my.id/api/

Environment Variables (in server/.env):
```
ALLOWED_ORIGINS=https://prototype-resort-apps.okiabrian.my.id
PORT=8084
GIN_MODE=release
```

### Service Token Configuration

For securing communication between frontend and backend:

1. Create a service token in Zero Trust > Access > Service Auth
2. Name it "Resort Apps Backend"
3. Configure environment variables in the backend:
   ```
   CF_ACCESS_CLIENT_ID=[Client ID from Cloudflare]
   CF_ACCESS_CLIENT_SECRET=[Client Secret from Cloudflare]
   CF_TEAM_DOMAIN=https://[your-team-name].cloudflareaccess.com
   CF_POLICY_AUD=[Application Audience Tag from Cloudflare]
   ```

### Testing

After configuration:

1. Visit https://prototype-resort-apps.okiabrian.my.id
2. You should be redirected to Cloudflare Access authentication
3. After successful authentication, you should see the application
4. Test API endpoints to ensure they work correctly

### Troubleshooting

Common issues:
- **CORS errors**: Ensure ALLOWED_ORIGINS matches your domain
- **Authentication loops**: Check identity provider configuration
- **Mixed content warnings**: Ensure all resources use HTTPS