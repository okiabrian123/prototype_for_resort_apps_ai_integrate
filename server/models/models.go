package models

import "time"

// Booking represents a booking entity with payment information
type Booking struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	ResortName   string    `json:"resort_name"`
	CheckIn      string    `json:"check_in"`
	CheckOut     string    `json:"check_out"`
	Guests       int       `json:"guests"`
	TotalPrice   float64   `json:"total_price"`
	Status       string    `json:"status"` // pending, confirmed, paid, cancelled
	PaymentDate  string    `json:"payment_date,omitempty"`
	CustomerName string    `json:"customer_name"`
	PhoneNumber  string    `json:"phone_number"`
	CreatedAt    time.Time `json:"created_at"`
}
