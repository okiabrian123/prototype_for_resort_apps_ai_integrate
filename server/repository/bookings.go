package repository

import (
	"database/sql"
	"log"
	"resort-app-server/database"
	"resort-app-server/models"
)

// GetAllBookings retrieves all bookings from the database
func GetAllBookings() ([]models.Booking, error) {
	rows, err := database.DB.Query("SELECT id, user_id, resort_name, check_in, check_out, guests, total_price, status, payment_date, customer_name, phone_number, created_at FROM bookings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		var customerName, phoneNumber sql.NullString
		err := rows.Scan(&booking.ID, &booking.UserID, &booking.ResortName, &booking.CheckIn, &booking.CheckOut, &booking.Guests, &booking.TotalPrice, &booking.Status, &booking.PaymentDate, &customerName, &phoneNumber, &booking.CreatedAt)
		if err != nil {
			log.Println("Error scanning booking row:", err)
			continue
		}
		// Handle NULL values
		if customerName.Valid {
			booking.CustomerName = customerName.String
		}
		if phoneNumber.Valid {
			booking.PhoneNumber = phoneNumber.String
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

// GetBookingByID retrieves a booking by its ID
func GetBookingByID(id int) (*models.Booking, error) {
	var booking models.Booking
	var customerName, phoneNumber sql.NullString
	err := database.DB.QueryRow("SELECT id, user_id, resort_name, check_in, check_out, guests, total_price, status, payment_date, customer_name, phone_number, created_at FROM bookings WHERE id = ?", id).
		Scan(&booking.ID, &booking.UserID, &booking.ResortName, &booking.CheckIn, &booking.CheckOut, &booking.Guests, &booking.TotalPrice, &booking.Status, &booking.PaymentDate, &customerName, &phoneNumber, &booking.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Handle NULL values
	if customerName.Valid {
		booking.CustomerName = customerName.String
	}
	if phoneNumber.Valid {
		booking.PhoneNumber = phoneNumber.String
	}

	return &booking, nil
}

// CreateBooking inserts a new booking into the database
func CreateBooking(booking *models.Booking) error {
	result, err := database.DB.Exec(
		"INSERT INTO bookings (user_id, resort_name, check_in, check_out, guests, total_price, status, payment_date, customer_name, phone_number) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		booking.UserID, booking.ResortName, booking.CheckIn, booking.CheckOut, booking.Guests, booking.TotalPrice, booking.Status, booking.PaymentDate, booking.CustomerName, booking.PhoneNumber)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	booking.ID = int(id)
	return nil
}

// UpdateBooking updates an existing booking in the database
func UpdateBooking(booking *models.Booking) error {
	_, err := database.DB.Exec(
		"UPDATE bookings SET user_id = ?, resort_name = ?, check_in = ?, check_out = ?, guests = ?, total_price = ?, status = ?, payment_date = ?, customer_name = ?, phone_number = ? WHERE id = ?",
		booking.UserID, booking.ResortName, booking.CheckIn, booking.CheckOut, booking.Guests, booking.TotalPrice, booking.Status, booking.PaymentDate, booking.CustomerName, booking.PhoneNumber, booking.ID)

	return err
}

// DeleteBooking removes a booking from the database
func DeleteBooking(id int) error {
	_, err := database.DB.Exec("DELETE FROM bookings WHERE id = ?", id)
	return err
}

// GetBookingsByStatus retrieves bookings by their status
func GetBookingsByStatus(status string) ([]models.Booking, error) {
	rows, err := database.DB.Query("SELECT id, user_id, resort_name, check_in, check_out, guests, total_price, status, payment_date, customer_name, phone_number, created_at FROM bookings WHERE status = ?", status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		var customerName, phoneNumber sql.NullString
		err := rows.Scan(&booking.ID, &booking.UserID, &booking.ResortName, &booking.CheckIn, &booking.CheckOut, &booking.Guests, &booking.TotalPrice, &booking.Status, &booking.PaymentDate, &customerName, &phoneNumber, &booking.CreatedAt)
		if err != nil {
			log.Println("Error scanning booking row:", err)
			continue
		}
		// Handle NULL values
		if customerName.Valid {
			booking.CustomerName = customerName.String
		}
		if phoneNumber.Valid {
			booking.PhoneNumber = phoneNumber.String
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

// GetBookingsByUserID retrieves bookings by user ID
func GetBookingsByUserID(userID int) ([]models.Booking, error) {
	rows, err := database.DB.Query("SELECT id, user_id, resort_name, check_in, check_out, guests, total_price, status, payment_date, customer_name, phone_number, created_at FROM bookings WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		var customerName, phoneNumber sql.NullString
		err := rows.Scan(&booking.ID, &booking.UserID, &booking.ResortName, &booking.CheckIn, &booking.CheckOut, &booking.Guests, &booking.TotalPrice, &booking.Status, &booking.PaymentDate, &customerName, &phoneNumber, &booking.CreatedAt)
		if err != nil {
			log.Println("Error scanning booking row:", err)
			continue
		}
		// Handle NULL values
		if customerName.Valid {
			booking.CustomerName = customerName.String
		}
		if phoneNumber.Valid {
			booking.PhoneNumber = phoneNumber.String
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

// GetBookingsByCustomerInfo retrieves bookings by customer name and phone number
// This is used for anonymous booking systems where customers don't have accounts
func GetBookingsByCustomerInfo(name, phone string) ([]models.Booking, error) {
	rows, err := database.DB.Query("SELECT id, user_id, resort_name, check_in, check_out, guests, total_price, status, payment_date, customer_name, phone_number, created_at FROM bookings WHERE customer_name = ? AND phone_number = ?", name, phone)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		var customerName, phoneNumber sql.NullString
		err := rows.Scan(&booking.ID, &booking.UserID, &booking.ResortName, &booking.CheckIn, &booking.CheckOut, &booking.Guests, &booking.TotalPrice, &booking.Status, &booking.PaymentDate, &customerName, &phoneNumber, &booking.CreatedAt)
		if err != nil {
			log.Println("Error scanning booking row:", err)
			continue
		}
		// Handle NULL values
		if customerName.Valid {
			booking.CustomerName = customerName.String
		}
		if phoneNumber.Valid {
			booking.PhoneNumber = phoneNumber.String
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}