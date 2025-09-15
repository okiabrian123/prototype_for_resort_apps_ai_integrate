package config

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/coreos/go-oidc/v3/oidc"
)

// CloudflareZeroTrustConfig holds the configuration for Cloudflare Zero Trust integration
type CloudflareZeroTrustConfig struct {
	TeamDomain string
	PolicyAUD  string
	Enabled    bool
}

// NewCloudflareZeroTrustConfig creates a new Cloudflare Zero Trust configuration
func NewCloudflareZeroTrustConfig() *CloudflareZeroTrustConfig {
	return &CloudflareZeroTrustConfig{
		TeamDomain: os.Getenv("CF_TEAM_DOMAIN"),
		PolicyAUD:  os.Getenv("CF_POLICY_AUD"),
		Enabled:    os.Getenv("CF_ZERO_TRUST_ENABLED") == "true",
	}
}

// ValidateCloudflareAccessJWT returns a Gin middleware for validating Cloudflare Access JWT tokens
func (cfg *CloudflareZeroTrustConfig) ValidateCloudflareAccessJWT() gin.HandlerFunc {
	// If Zero Trust is not enabled, return a no-op middleware
	if !cfg.Enabled {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	// Set default team domain if not provided
	teamDomain := cfg.TeamDomain
	if teamDomain == "" {
		teamDomain = "https://your-team.cloudflareaccess.com" // Replace with your actual team domain
	}

	// Configure OIDC verifier
	ctx := context.Background()
	certsURL := teamDomain + "/cdn-cgi/access/certs"
	
	config := &oidc.Config{
		ClientID: cfg.PolicyAUD,
	}
	
	keySet := oidc.NewRemoteKeySet(ctx, certsURL)
	verifier := oidc.NewVerifier(teamDomain, keySet, config)

	return func(c *gin.Context) {
		// Skip validation for health check and root endpoints
		if c.Request.URL.Path == "/health" || c.Request.URL.Path == "/" {
			c.Next()
			return
		}

		// Get the CF Authorization token from headers
		accessJWT := c.GetHeader("Cf-Access-Jwt-Assertion")
		if accessJWT == "" {
			// Also check for cookie (for browser requests)
			if cookie, err := c.Request.Cookie("CF_Authorization"); err == nil {
				accessJWT = cookie.Value
			}
		}

		if accessJWT == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing Cloudflare Access token",
			})
			c.Abort()
			return
		}

		// Verify the access token
		_, err := verifier.Verify(c.Request.Context(), accessJWT)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Cloudflare Access token: " + err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}