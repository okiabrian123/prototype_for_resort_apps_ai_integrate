package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
	"resort-app-server/models"
	"strings"
)

// GetHouses retrieves all houses from the JSON file
func GetHouses() ([]models.House, error) {
	// Get the absolute path to the resorts.json file
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(currentDir, "data", "houses.json")

	// Read the JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data
	var houses []models.House
	err = json.Unmarshal(data, &houses)
	if err != nil {
		return nil, err
	}

	return houses, nil
}

// GetHouseByID retrieves a specific house by its ID
func GetHouseByID(id int) (*models.House, error) {
	houses, err := GetHouses()
	if err != nil {
		return nil, err
	}

	for _, house := range houses {
		if house.ID == id {
			return &house, nil
		}
	}

	return nil, nil
}

// SearchHouses searches for houses by name or location
func SearchHouses(query string) ([]models.House, error) {
	houses, err := GetHouses()
	if err != nil {
		return nil, err
	}

	var results []models.House
	queryLower := strings.ToLower(query)

	for _, house := range houses {
		// Check if query matches house name or location
		nameLower := strings.ToLower(house.Name)
		locationLower := strings.ToLower(house.Location)

		if strings.Contains(nameLower, queryLower) || strings.Contains(locationLower, queryLower) {
			results = append(results, house)
		}
	}

	return results, nil
}

// GetHousesByGuests returns houses that can accommodate at least the specified number of guests
func GetHousesByGuests(guests int) ([]models.House, error) {
	houses, err := GetHouses()
	if err != nil {
		return nil, err
	}

	var filteredHouses []models.House
	for _, house := range houses {
		if house.Guests >= guests {
			filteredHouses = append(filteredHouses, house)
		}
	}

	return filteredHouses, nil
}
