package main

import (
	"net/http"
	"strconv"

	"resort-app-server/models"
	"resort-app-server/repository"

	"github.com/gin-gonic/gin"
)

// getHouses returns all houses
func getHouses(c *gin.Context) {
	houses, err := repository.GetHouses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve houses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"houses": houses,
		"count":  len(houses),
	})
}

// getHousesByGuests returns houses that can accommodate the specified number of guests
func getHousesByGuests(c *gin.Context) {
	// Get the guests parameter from query
	guestsParam := c.Query("guests")
	if guestsParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guests parameter is required"})
		return
	}

	// Convert guests parameter to integer
	guests, err := strconv.Atoi(guestsParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guests parameter must be a valid number"})
		return
	}

	// Get all houses
	houses, err := repository.GetHouses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve houses"})
		return
	}

	// Filter houses based on guest capacity
	var filteredHouses []models.House
	for _, house := range houses {
		if house.Guests >= guests {
			filteredHouses = append(filteredHouses, house)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"houses": filteredHouses,
		"count":  len(filteredHouses),
	})
}

// getHouse returns a specific house by ID
func getHouse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid house ID"})
		return
	}

	house, err := repository.GetHouseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve house"})
		return
	}

	if house == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "House not found"})
		return
	}

	c.JSON(http.StatusOK, house)
}

// searchHouses searches for houses by name or location
func searchHouses(c *gin.Context) {
	query := c.Param("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	houses, err := repository.SearchHouses(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search houses"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"houses": houses,
		"count":  len(houses),
		"query":  query,
	})
}
