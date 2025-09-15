# DNS Configuration for prototype-resort-apps.okiabrian.my.id

## DNS Records Setup

To properly configure your domain with Cloudflare, you'll need to set up the following DNS records:

### Primary A Record
```
Type: A
Name: prototype-resort-apps
Content: [YOUR_SERVER_PUBLIC_IP]
Proxy status: Proxied
TTL: Auto
```

This record points your subdomain to your server's IP address and enables Cloudflare's proxy features.

### Optional WWW Record
```
Type: CNAME
Name: www.prototype-resort-apps
Content: prototype-resort-apps.okiabrian.my.id
Proxy status: Proxied
TTL: Auto
```

This record allows users to access your site via www.prototype-resort-apps.okiabrian.my.id.

## Cloudflare Nameserver Configuration

After adding your site to Cloudflare, you'll need to update your domain's nameservers at your domain registrar:

1. In Cloudflare dashboard, go to your site's Overview page
2. Scroll down to find your assigned Cloudflare nameservers (usually 2-4 ns.cloudflare.com addresses)
3. Log in to your domain registrar's control panel (where you purchased okiabrian.my.id)
4. Update the nameservers to the ones provided by Cloudflare
5. Wait for DNS propagation (this can take up to 24 hours, but is usually much faster)

## DNS Validation

To verify your DNS configuration is working correctly:

1. Use `dig` command:
   ```bash
   dig prototype-resort-apps.okiabrian.my.id
   ```

2. Or use online tools like:
   - https://dnschecker.org
   - https://www.whatsmydns.net

You should see your Cloudflare nameservers in the response.

## Health Checks

Cloudflare can monitor your server's availability:

1. Go to Traffic > Health Checks in Cloudflare dashboard
2. Create a new health check:
   - Name: Resort Apps Health Check
   - Host: prototype-resort-apps.okiabrian.my.id
   - Path: /health
   - Expected status: 200
   - Check interval: 60 seconds

This will monitor your backend's health endpoint that's already implemented in your Go server.

## Load Balancing (Optional)

If you plan to deploy multiple instances of your application:

1. Go to Traffic > Load Balancing
2. Create a new load balancer:
   - Name: Resort Apps Load Balancer
   - Hostname: prototype-resort-apps.okiabrian.my.id
   - Add your origin servers
   - Configure health checks
   - Set load balancing method (e.g., Round Robin)

## DNS Security

Enable these DNS security features in Cloudflare:

1. Go to DNS > Settings
2. Enable:
   - DNSSEC: On
   - Email obfuscation: On
   - Hotlink protection: On
   - Server side exclude: On
   - Rocket Loader: On (for better performance)

## Caching Configuration

For optimal performance:

1. Go to Speed > Optimization
2. Enable:
   - Auto Minify: JavaScript, CSS, HTML
   - Brotli compression: On
   - Rocket Loader: On

3. In Caching > Configuration:
   - Caching level: Standard
   - Browser Cache TTL: 4 hours
   - Always Online: On

## Custom Error Pages

Create custom error pages for better user experience:

1. Go to Custom Pages in Cloudflare dashboard
2. Customize:
   - Always Online
   - 1000, 1001, 1002, 1003, 1004 error pages
   - 500, 501, 502, 503, 504, 505 error pages

## Monitoring

Set up monitoring for your DNS records:

1. Use Cloudflare's analytics to monitor:
   - DNS queries
   - Cached vs uncached requests
   - Threats blocked
   - Bandwidth usage

2. Set up alerts:
   - Certificate transparency alerts
   - Origin monitoring alerts
   - Health check alerts

This configuration will ensure your domain is properly set up with Cloudflare's DNS services, providing performance benefits and security features.