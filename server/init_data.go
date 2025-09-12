package main

import (
	"log"
	"resort-app-server/models"
	"resort-app-server/repository"
)

// initSampleData initializes the database with sample data
func initSampleData() {
	// Check if we already have data
	bookings, err := repository.GetAllBookings()
	if err != nil {
		log.Fatal("Failed to check existing bookings:", err)
	}

	// If no bookings exist, add sample data
	if len(bookings) == 0 {
		log.Println("Initializing database with sample data...")

		// Sample bookings with different statuses
		sampleBookings := []models.Booking{
			{
				UserID:       101,
				ResortName:   "Paradise Beach Resort",
				CheckIn:      "2025-10-15",
				CheckOut:     "2025-10-20",
				Guests:       2,
				TotalPrice:   1200.00,
				Status:       "paid",
				CustomerName: "John Smith",
				PhoneNumber:  "+1234567890",
			},
			{
				UserID:       102,
				ResortName:   "Mountain View Lodge",
				CheckIn:      "2025-11-05",
				CheckOut:     "2025-11-10",
				Guests:       4,
				TotalPrice:   2100.00,
				Status:       "confirmed",
				CustomerName: "Jane Doe",
				PhoneNumber:  "+0987654321",
			},
			{
				UserID:       103,
				ResortName:   "Urban Luxury Hotel",
				CheckIn:      "2025-09-20",
				CheckOut:     "2025-09-25",
				Guests:       2,
				TotalPrice:   1500.00,
				Status:       "pending",
				CustomerName: "Bob Johnson",
				PhoneNumber:  "+1122334455",
			},
		}

		// Insert sample bookings
		for i := range sampleBookings {
			err := repository.CreateBooking(&sampleBookings[i])
			if err != nil {
				log.Printf("Failed to create booking: %v", err)
			} else {
				log.Printf("Created booking ID: %d", sampleBookings[i].ID)
			}
		}

		log.Println("Sample data initialization completed")
	} else {
		log.Println("Database already contains data, skipping initialization")
	}
}
