# SSL/TLS Configuration for prototype-resort-apps.okiabrian.my.id

## Overview

This document details how to configure SSL/TLS encryption for your domain using Cloudflare's certificate management. Cloudflare provides free SSL/TLS certificates that automatically renew, ensuring your application is always secured with HTTPS.

## SSL/TLS Encryption Modes

Cloudflare offers several encryption modes. For maximum security with your prototype application, use "Full (strict)" mode.

### Encryption Mode Options

1. **Off**: No SSL/TLS encryption (Not recommended)
2. **Flexible**: Encrypts traffic between browser and Cloudflare only
3. **Full**: Encrypts traffic between browser and Cloudflare, and between Cloudflare and origin server
4. **Full (strict)**: Same as Full, but requires valid SSL certificate on origin server

## Step 1: Configure SSL/TLS Encryption Mode

1. Log in to Cloudflare dashboard
2. Select your domain (okiabrian.my.id)
3. Go to SSL/TLS > Overview
4. Set encryption mode to "Full (strict)"
5. This ensures:
   - End-to-end encryption
   - Certificate validation on your origin server
   - Maximum security for data in transit

## Step 2: Configure Edge Certificates

1. Go to SSL/TLS > Edge Certificates
2. Configure these settings:

### Certificate Status and Transparency
- Always Use HTTPS: On
- Minimum TLS Version: TLS 1.2
- Opportunistic Encryption: On
- TLS 1.3: On
- Automatic HTTPS Rewrites: On

### Certificate Validation

With "Full (strict)" mode, you need a valid SSL certificate on your origin server:

#### Option 1: Use Origin CA (Recommended)
1. Go to SSL/TLS > Origin Server
2. Click "Create Certificate"
3. Configure:
   - Certificate hostnames: prototype-resort-apps.okiabrian.my.id
   - Certificate validity: 15 years (maximum)
4. Click "Create"
5. Save the certificate and private key to your server
6. Configure your web server (nginx, Apache, etc.) to use these certificates

#### Option 2: Use Let's Encrypt
1. Install Certbot on your server:
   ```bash
   sudo apt-get install certbot
   ```
2. Generate certificate:
   ```bash
   sudo certbot certonly --standalone -d prototype-resort-apps.okiabrian.my.id
   ```
3. Configure your web server to use the generated certificates

## Step 3: Configure Custom SSL (If Needed)

If you have your own SSL certificate:

1. Go to SSL/TLS > Custom SSL
2. Click "Upload Custom SSL Certificate"
3. Paste your certificate and private key
4. Select the appropriate bundle method
5. Click "Upload"

## Step 4: Configure HSTS (HTTP Strict Transport Security)

1. Go to SSL/TLS > Edge Certificates
2. Enable HSTS with these settings:
   - HSTS: On
   - Max-age: 12 months (31536000 seconds)
   - Include subdomains: On
   - Preload: Off (unless you're submitting to browser preload lists)

## Step 5: Configure Authenticated Origin Pulls

To ensure requests to your origin server come from Cloudflare:

1. Go to SSL/TLS > Origin Server
2. Enable "Authenticated Origin Pulls"
3. This requires:
   - Valid certificate on your origin server
   - Configuration to validate Cloudflare's certificate

## Step 6: Configure Browser Compatibility

Ensure compatibility with older browsers:

1. Go to SSL/TLS > Edge Certificates
2. Set Minimum TLS Version to TLS 1.2 (this is the recommended minimum for security)
3. Older versions have known vulnerabilities

## Step 7: Mixed Content Fixes

Cloudflare can automatically fix mixed content issues:

1. Go to SSL/TLS > Edge Certificates
2. Enable "Automatic HTTPS Rewrites"
3. This automatically rewrites HTTP links to HTTPS in your HTML content

## Testing SSL/TLS Configuration

### Online Testing Tools

Use these tools to verify your SSL/TLS configuration:

1. SSL Labs SSL Test: https://www.ssllabs.com/ssltest/
2. Cloudflare SSL/TLS Analyzer
3. Mozilla Observatory: https://observatory.mozilla.org/

### Command Line Testing

```bash
# Test certificate chain
openssl s_client -connect prototype-resort-apps.okiabrian.my.id:443 -servername prototype-resort-apps.okiabrian.my.id

# Test specific TLS versions
openssl s_client -connect prototype-resort-apps.okiabrian.my.id:443 -servername prototype-resort-apps.okiabrian.my.id -tls1_2
```

## Certificate Management

### Automatic Renewal

Cloudflare certificates automatically renew, but for Origin CA or Let's Encrypt:

1. Set up automatic renewal scripts
2. Monitor certificate expiration dates
3. Test renewal process regularly

### Certificate Monitoring

1. Use Cloudflare's certificate monitoring
2. Set up alerts for certificate expiration
3. Regularly check certificate transparency logs

## Security Headers

In addition to SSL/TLS, configure these security headers in Cloudflare:

1. Go to Rules > Transform Rules
2. Create a new transform rule
3. Add these response headers:
   ```
   Strict-Transport-Security: max-age=31536000; includeSubDomains; preload
   X-Content-Type-Options: nosniff
   X-Frame-Options: DENY
   X-XSS-Protection: 1; mode=block
   Content-Security-Policy: default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'
   Referrer-Policy: strict-origin-when-cross-origin
   Permissions-Policy: geolocation=(), microphone=(), camera=()
   ```

## Troubleshooting Common Issues

### Certificate Not Trusted

If browsers show certificate warnings:

1. Verify encryption mode is set correctly
2. Check that origin server has valid certificate (for Full/Full(strict) modes)
3. Ensure DNS records are properly proxied through Cloudflare

### Mixed Content Warnings

If you see mixed content warnings:

1. Enable "Automatic HTTPS Rewrites"
2. Check your application code for hardcoded HTTP URLs
3. Use relative URLs where possible
4. Implement Content Security Policy headers

### SSL Handshake Errors

For SSL handshake errors:

1. Check Minimum TLS Version compatibility
2. Verify cipher suite compatibility
3. Ensure certificate chain is complete
4. Check for firewall/proxy interference

### Performance Optimization

To optimize SSL/TLS performance:

1. Enable HTTP/2 in Cloudflare settings
2. Use Cloudflare's Argo Smart Routing
3. Enable Brotli compression
4. Configure appropriate cache settings