package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the SQLite database
func InitDB() {
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		// Try to load from current directory as fallback
		err = godotenv.Load(".env")
		if err != nil {
			log.Println("No .env file found, using default settings")
		}
	}

	// Create data directory if it doesn't exist
	dataDir := "data"
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.Mkdir(dataDir, 0755)
	}

	// Get database path from environment variable or use default
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = filepath.Join(dataDir, "resort.db")
	}

	// Open database connection
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Create tables
	createTables()
}

// createTables creates the necessary tables if they don't exist
func createTables() {
	// Create bookings table with payment status
	bookingsTable := `
	CREATE TABLE IF NOT EXISTS bookings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		resort_name TEXT NOT NULL,
		check_in DATE NOT NULL,
		check_out DATE NOT NULL,
		guests INTEGER NOT NULL,
		total_price REAL NOT NULL,
		status TEXT NOT NULL DEFAULT 'pending', -- pending, confirmed, paid, cancelled
		payment_date DATE,
		customer_name TEXT,
		phone_number TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(bookingsTable)
	if err != nil {
		log.Fatal("Failed to create bookings table:", err)
	}

	log.Println("Database tables created successfully")
}
