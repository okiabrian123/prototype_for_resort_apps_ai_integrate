#!/bin/bash

# Cloudflare Tunnel Setup Script for Resort Apps
# This script helps automate the setup of Cloudflare Tunnel for the prototype

set -e  # Exit on any error

# Load environment variables from .env if it exists
load_env_vars() {
    local env_file="$PWD/.env"
    if [ -f "$env_file" ]; then
        echo "Loading environment variables from $env_file"
        export $(grep -v '^#' "$env_file" | xargs)
    else
        echo "Error: Environment file $env_file not found. Please create .env file with required configuration."
        exit 1
    fi
    
    # Verify that required environment variables are set
    if [ -z "$PORT" ] || [ -z "$FRONTEND_PORT" ] || [ -z "$DOMAIN_NAME" ] || [ -z "$TUNNEL_NAME" ] || [ -z "$TUNNEL_UUID" ]; then
        echo "Error: PORT, FRONTEND_PORT, DOMAIN_NAME, TUNNEL_NAME, and TUNNEL_UUID must be defined in .env file"
        exit 1
    fi
}

# Function to check if cloudflared is installed
check_cloudflared() {
    if ! command -v cloudflared &> /dev/null
    then
        echo "cloudflared could not be found. Please install it first."
        echo "Visit: https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/downloads/"
        exit 1
    fi
    echo "✓ cloudflared is installed"
}

# Function to run the tunnel with automatic restart
run_tunnel() {
    # Load environment variables
    load_env_vars
    
    local config_file="$HOME/.cloudflared/${TUNNEL_NAME}.yml"
    
    echo "✓ Using tunnel name: $TUNNEL_NAME"
    echo "✓ Using tunnel UUID: $TUNNEL_UUID"
    echo "✓ Using backend port: $PORT"
    echo "✓ Using frontend port: $FRONTEND_PORT"
    
    # Update configuration file with latest settings
    echo "Updating configuration file..."
    CONFIG_DIR="$HOME/.cloudflared"
    
    mkdir -p $CONFIG_DIR
    
    cat > $config_file << EOF
tunnel: $TUNNEL_NAME
credentials-file: $CONFIG_DIR/$TUNNEL_UUID.json

# Ingress rules for routing traffic
ingress:
  # Backend API
  - hostname: $DOMAIN_NAME
    path: /api/
    service: http://localhost:$PORT
  # Backend health endpoint
  - hostname: $DOMAIN_NAME
    path: /health
    service: http://localhost:$PORT
  # Frontend application (Vite dev server) - catch-all for all other paths
  - hostname: $DOMAIN_NAME
    service: http://localhost:$FRONTEND_PORT
  # Catch-all rule
  - service: http_status:404
EOF

    echo "✓ Configuration file updated at $config_file"
    
    echo "Starting Cloudflare Tunnel with automatic restart..."
    echo "Configuration file: $config_file"
    echo "Tunnel name: $TUNNEL_NAME"
    echo "Press Ctrl+C to stop"
    
    # Run tunnel with automatic restart
    while true; do
        echo "Starting tunnel at $(date)"
        cloudflared tunnel --config "$config_file" run "$TUNNEL_NAME"
        
        # If the tunnel exits, wait 5 seconds before restarting
        echo "Tunnel stopped at $(date). Restarting in 5 seconds..."
        sleep 5
    done
}

# Function to stop any running tunnel instances
stop_tunnel() {
    echo "Stopping any running Cloudflare Tunnel instances..."
    pkill -f "cloudflared tunnel" || true
    echo "✓ Tunnel instances stopped"
}

# Function to show tunnel status
tunnel_status() {
    echo "Checking Cloudflare Tunnel status..."
    if pgrep -f "cloudflared tunnel" > /dev/null; then
        echo "✓ Cloudflare Tunnel is running"
        pgrep -f "cloudflared tunnel"
    else
        echo "✗ Cloudflare Tunnel is not running"
    fi
}

# Main setup function
setup_tunnel() {
    echo "Cloudflare Tunnel Setup for Resort Apps"
    echo "========================================"

    # Load environment variables
    load_env_vars

    # Check if cloudflared is installed
    check_cloudflared

    # Authenticate with Cloudflare
    echo "Step 1: Authenticating with Cloudflare..."
    echo "A browser window will open. Please log in to your Cloudflare account."
    read -p "Press Enter to continue..."
    cloudflared tunnel login

    # Create tunnel
    echo "Step 2: Creating tunnel..."
    # Use values from environment
    TUNNEL_NAME_ENV="$TUNNEL_NAME"
    TUNNEL_UUID_ENV="$TUNNEL_UUID"
    
    # Check if tunnel already exists
    if cloudflared tunnel info $TUNNEL_NAME_ENV >/dev/null 2>&1; then
        echo "✓ Tunnel '$TUNNEL_NAME_ENV' already exists"
    else
        # Create tunnel with specific UUID
        if cloudflared tunnel create --uuid $TUNNEL_UUID_ENV $TUNNEL_NAME_ENV >/dev/null 2>&1; then
            echo "✓ Tunnel created with UUID: $TUNNEL_UUID_ENV"
        else
            echo "Warning: Could not create tunnel with specific UUID, creating with random UUID"
            cloudflared tunnel create $TUNNEL_NAME_ENV
            # Get the actual UUID
            TUNNEL_INFO=$(cloudflared tunnel info $TUNNEL_NAME_ENV)
            TUNNEL_UUID_ENV=$(echo "$TUNNEL_INFO" | grep -o '[0-9a-f]\{8\}-[0-9a-f]\{4\}-[0-9a-f]\{4\}-[0-9a-f]\{4\}-[0-9a-f]\{12\}' | head -1)
        fi
    fi

    echo "✓ Using tunnel UUID: $TUNNEL_UUID_ENV"

    # Create configuration file
    echo "Step 4: Creating configuration file..."
    CONFIG_DIR="$HOME/.cloudflared"
    CONFIG_FILE="$CONFIG_DIR/${TUNNEL_NAME_ENV}.yml"

    mkdir -p $CONFIG_DIR

    cat > $CONFIG_FILE << EOF
tunnel: $TUNNEL_NAME_ENV
credentials-file: $CONFIG_DIR/$TUNNEL_UUID_ENV.json

# Ingress rules for routing traffic
ingress:
  # Backend API
  - hostname: $DOMAIN_NAME
    path: /api/
    service: http://localhost:$PORT
  # Backend health endpoint
  - hostname: $DOMAIN_NAME
    path: /health
    service: http://localhost:$PORT
  # Frontend application (Vite dev server) - catch-all for all other paths
  - hostname: $DOMAIN_NAME
    service: http://localhost:$FRONTEND_PORT
  # Catch-all rule
  - service: http_status:404
EOF

    echo "✓ Configuration file created at $CONFIG_FILE"

    # Route DNS
    echo "Step 5: Routing DNS..."
    cloudflared tunnel route dns $TUNNEL_NAME_ENV $DOMAIN_NAME

    echo "✓ DNS routed to tunnel"

    echo ""
    echo "Setup complete!"
    echo "==============="
    echo "Next steps:"
    echo "1. Ensure your frontend (port $FRONTEND_PORT) and backend (port $PORT) applications are running"
    echo "2. Run the tunnel with: ./setup-cloudflare-tunnel.sh run"
    echo "3. Visit https://$DOMAIN_NAME"
    echo "4. Configure Zero Trust Application in the Cloudflare dashboard"
}

# Parse command line arguments
case "$1" in
    run)
        run_tunnel
        ;;
    stop)
        stop_tunnel
        ;;
    status)
        tunnel_status
        ;;
    setup)
        setup_tunnel
        ;;
    *)
        echo "Cloudflare Tunnel Manager for Resort Apps"
        echo "Usage: $0 {setup|run|stop|status}"
        echo ""
        echo "Commands:"
        echo "  setup   - Set up the Cloudflare Tunnel (authentication, creation, configuration)"
        echo "  run     - Run the tunnel with automatic restart"
        echo "  stop    - Stop any running tunnel instances"
        echo "  status  - Check tunnel status"
        echo ""
        echo "Examples:"
        echo "  $0 setup   # Set up the tunnel"
        echo "  $0 run     # Run the tunnel with automatic restart"
        echo "  $0 stop    # Stop the tunnel"
        echo "  $0 status  # Check tunnel status"
        ;;
esac