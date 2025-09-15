package function_calling

import (
	"encoding/json"
	"fmt"

	"resort-app-server/models"
	"resort-app-server/repository"
)

// HouseListData represents the data structure for house list function calling
type HouseListData struct {
	Guests int `json:"guests"`
}

// HouseOption represents a house option for the frontend
type HouseOption struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Guests        int     `json:"guests"`
	PricePerNight float64 `json:"price_per_night"`
	ImageURL      string  `json:"image_url"`
}

// ExtractHouseListData extracts house list data from AI response when function calling is executed
func ExtractHouseListData(messageContent string) (*HouseListData, bool) {
	// Look for house list data in the format specified in the system prompt
	startTag := "<HOUSE_LIST_DATA>"
	endTag := "</HOUSE_LIST_DATA>"

	startIdx := FindStringIndex(messageContent, startTag)
	if startIdx == -1 {
		return nil, false
	}

	endIdx := FindStringIndex(messageContent, endTag)
	if endIdx == -1 {
		return nil, false
	}

	// Extract the JSON data between the tags
	houseListDataJSON := messageContent[startIdx+len(startTag) : endIdx]
	// Clean up whitespace
	houseListDataJSON = TrimString(houseListDataJSON)

	// Parse the JSON into a HouseListData struct
	var houseListData HouseListData
	if err := json.Unmarshal([]byte(houseListDataJSON), &houseListData); err != nil {
		fmt.Printf("Error parsing house list data: %v\n", err)
		return nil, false
	}

	return &houseListData, true
}

// GetHousesByGuestsForAI retrieves houses based on guest count for AI function calling
func GetHousesByGuestsForAI(guests int) ([]models.House, error) {
	return repository.GetHousesByGuests(guests)
}

// ProcessHouseListData processes house list data and returns a structured response
func ProcessHouseListData(houseListData *HouseListData) (interface{}, error) {
	// Get houses based on guest count
	houses, err := GetHousesByGuestsForAI(houseListData.Guests)
	if err != nil {
		return nil, fmt.Errorf("error retrieving houses: %v", err)
	}

	if len(houses) > 0 {
		// Sort houses: first by exact guest count match, then by price
		// We'll do this sorting manually to avoid importing additional packages
		for i := 0; i < len(houses)-1; i++ {
			for j := i + 1; j < len(houses); j++ {
				// Prioritize houses that exactly match the requested guest count
				exactMatchI := houses[i].Guests == houseListData.Guests
				exactMatchJ := houses[j].Guests == houseListData.Guests

				shouldSwap := false
				if exactMatchI && !exactMatchJ {
					// i should come first, no swap needed
					shouldSwap = false
				} else if !exactMatchI && exactMatchJ {
					// j should come first, swap needed
					shouldSwap = true
				} else {
					// If both have the same match status, sort by price (ascending)
					shouldSwap = houses[i].PricePerNight > houses[j].PricePerNight
				}

				if shouldSwap {
					houses[i], houses[j] = houses[j], houses[i]
				}
			}
		}

		// Create a structured response that includes house details and image URLs
		var houseOptionsList []HouseOption
		for _, house := range houses {
			houseOptionsList = append(houseOptionsList, HouseOption{
				ID:            house.ID,
				Name:          house.Name,
				Guests:        house.Guests,
				PricePerNight: house.PricePerNight,
				ImageURL:      house.ImageURL,
			})
		}

		// Convert to JSON for the frontend
		houseOptionsJSON, err := json.Marshal(houseOptionsList)
		if err != nil {
			return nil, fmt.Errorf("error marshaling house options: %v", err)
		}

		// Return a structured response that the frontend can parse
		return map[string]interface{}{
			"type":    "house_options",
			"houses":  string(houseOptionsJSON),
			"message": "Please select one of these houses:",
		}, nil
	} else {
		houseOptions := "I'm sorry, but we don't have any houses available for " +
			fmt.Sprintf("%d guests at the moment. Would you like to try a different number of guests?",
				houseListData.Guests)

		return map[string]interface{}{
			"message": houseOptions,
		}, nil
	}
}
