package main

import (
	"log"
	"os"

	"resort-app-server/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
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

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

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
	router.Run(":" + port)
}
