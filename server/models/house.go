package models

// House represents a house entity
type House struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Location      string   `json:"location"`
	Rating        float64  `json:"rating"`
	PricePerNight float64  `json:"price_per_night"`
	ImageURL      string   `json:"image_url"`
	Amenities     []string `json:"amenities"`
	Guests        int      `json:"guests"`
}
