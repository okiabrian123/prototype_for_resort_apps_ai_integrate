package main

import (
	"log"
	"os"

	"resort-app-server/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from root .env file
	err := godotenv.Load("../.env")
	if err != nil {
		// Try to load from current directory as fallback
		err = godotenv.Load(".env")
		if err != nil {
			log.Println("No .env file found, using default settings")
		}
	}

	// Initialize database
	database.InitDB()
	defer database.DB.Close()

	// Initialize sample data
	initSampleData()

	// Set Gin to release mode in production
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	router := gin.Default()

	// Zero Trust Security: Configure CORS with strict policies
	allowedOrigin := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigin == "" {
		allowedOrigin = "https://okiabrian.my.id"
	}

	// Add security headers middleware
	router.Use(func(c *gin.Context) {
		// Zero Trust Security: Set security headers
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'")

		// Configure CORS with strict policies for Zero Trust
		// Allow both the production domain and localhost for development
		origin := c.Request.Header.Get("Origin")
		if origin == "http://localhost:8085" {
			c.Header("Access-Control-Allow-Origin", "http://localhost:8085")
		} else if origin == "" {
			// For requests without Origin header, use the configured allowed origin
			c.Header("Access-Control-Allow-Origin", allowedOrigin)
		} else {
			// For other origins, check if they match the allowed origin
			c.Header("Access-Control-Allow-Origin", allowedOrigin)
		}

		c.Header("Access-Control-Allow-Credentials", "false")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Max-Age", "86400") // 24 hours

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Resort App Server",
			"status":  "success",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	// House routes
	houses := router.Group("/api/houses")
	{
		houses.GET("/", getHouses)
		houses.GET("/guests", getHousesByGuests)
		houses.GET("/:id", getHouse)
		houses.GET("/search/:query", searchHouses)
	}

	// Booking routes
	booking := router.Group("/api/bookings")
	{
		booking.GET("/", getBookings)
		booking.GET("/:id", getBooking)
		booking.POST("/", createBooking)
		booking.PUT("/:id", updateBooking)
		booking.DELETE("/:id", deleteBooking)
		booking.GET("/status/:status", getBookingsByStatus)
		booking.GET("/user/:user_id", getBookingsByUser)
		booking.GET("/customer", getBookingsByCustomerInfo)
	}

	// Chatbot routes
	chat := router.Group("/api/chat")
	{
		chat.POST("/message", chatWithAI)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("Allowed CORS origin: %s", allowedOrigin)
	router.Run(":" + port)
}