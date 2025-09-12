package function_calling

import (
	"fmt"
)

// IsFunctionCallingExecuted checks if the AI response contains function calling instructions
// by looking for the <BOOKING_DATA> or <HOUSE_LIST_DATA> tag in the message content
func IsFunctionCallingExecuted(messageContent string) bool {
	// Check if the message content contains the BOOKING_DATA or HOUSE_LIST_DATA tag
	return FindStringIndex(messageContent, "<BOOKING_DATA>") != -1 ||
		FindStringIndex(messageContent, "<HOUSE_LIST_DATA>") != -1 ||
		FindStringIndex(messageContent, "<HOUSE-TYPE_DATA>") != -1
}

// ProcessFunctionCalling processes the appropriate function calling based on the AI response
func ProcessFunctionCalling(messageContent string) (interface{}, bool, error) {
	// Check for house list data first
	houseListData, houseListFound := ExtractHouseListData(messageContent)
	if houseListFound {
		response, err := ProcessHouseListData(houseListData)
		return response, true, err
	}

	// Check for booking data
	bookingData, bookingFound := ExtractBookingData(messageContent)
	if bookingFound {
		// Validate the booking data
		if err := ValidateBookingData(bookingData); err != nil {
			return nil, true, fmt.Errorf("invalid booking data: %v", err)
		}

		// Save booking to database
		if err := SaveBookingToDatabase(bookingData); err != nil {
			return nil, true, fmt.Errorf("failed to save booking: %v", err)
		}

		// Return a user-friendly message
		return map[string]string{
			"message": "Thank you! Your booking is now pending confirmation from our receptionist. We'll contact you shortly about the payment.",
		}, true, nil
	}

	return nil, false, nil
}
