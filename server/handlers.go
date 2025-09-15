package main

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"resort-app-server/models"
	"resort-app-server/repository"
	"resort-app-server/tool_calling/function_calling"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

// getBookings returns all bookings
func getBookings(c *gin.Context) {
	bookings, err := repository.GetAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bookings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bookings": bookings,
		"count":    len(bookings),
	})
}

// getBooking returns a specific booking by ID
func getBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	booking, err := repository.GetBookingByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve booking"})
		return
	}

	if booking == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, booking)
}

// createBooking creates a new booking
func createBooking(c *gin.Context) {
	var bookingInput struct {
		UserID     int     `json:"user_id" binding:"required"`
		ResortName string  `json:"resort_name" binding:"required"`
		CheckIn    string  `json:"check_in" binding:"required"`
		CheckOut   string  `json:"check_out" binding:"required"`
		Guests     int     `json:"guests" binding:"required"`
		TotalPrice float64 `json:"total_price" binding:"required"`
		Status     string  `json:"status"`
	}

	if err := c.BindJSON(&bookingInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Set default status if not provided
	if bookingInput.Status == "" {
		bookingInput.Status = "pending"
	}

	booking := &models.Booking{
		UserID:     bookingInput.UserID,
		ResortName: bookingInput.ResortName,
		CheckIn:    bookingInput.CheckIn,
		CheckOut:   bookingInput.CheckOut,
		Guests:     bookingInput.Guests,
		TotalPrice: bookingInput.TotalPrice,
		Status:     bookingInput.Status,
	}

	err := repository.CreateBooking(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusCreated, booking)
}

// updateBooking updates an existing booking
func updateBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	// Check if booking exists
	existingBooking, err := repository.GetBookingByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve booking"})
		return
	}

	if existingBooking == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	var bookingInput struct {
		UserID      int     `json:"user_id"`
		ResortName  string  `json:"resort_name"`
		CheckIn     string  `json:"check_in"`
		CheckOut    string  `json:"check_out"`
		Guests      int     `json:"guests"`
		TotalPrice  float64 `json:"total_price"`
		Status      string  `json:"status"`
		PaymentDate string  `json:"payment_date"`
	}

	if err := c.BindJSON(&bookingInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Update the booking with new values
	updatedBooking := &models.Booking{
		ID:          id,
		UserID:      bookingInput.UserID,
		ResortName:  bookingInput.ResortName,
		CheckIn:     bookingInput.CheckIn,
		CheckOut:    bookingInput.CheckOut,
		Guests:      bookingInput.Guests,
		TotalPrice:  bookingInput.TotalPrice,
		Status:      bookingInput.Status,
		PaymentDate: bookingInput.PaymentDate,
		CreatedAt:   existingBooking.CreatedAt,
	}

	err = repository.UpdateBooking(updatedBooking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
		return
	}

	c.JSON(http.StatusOK, updatedBooking)
}

// deleteBooking removes a booking
func deleteBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	// Check if booking exists
	booking, err := repository.GetBookingByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve booking"})
		return
	}

	if booking == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	err = repository.DeleteBooking(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}

// getBookingsByStatus returns bookings filtered by status
func getBookingsByStatus(c *gin.Context) {
	status := c.Param("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status parameter is required"})
		return
	}

	bookings, err := repository.GetBookingsByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bookings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bookings": bookings,
		"count":    len(bookings),
		"status":   status,
	})
}

// getBookingsByUser returns bookings for a specific user
func getBookingsByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	bookings, err := repository.GetBookingsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bookings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bookings": bookings,
		"count":    len(bookings),
		"user_id":  userID,
	})
}

// ChatMessage represents a message in the chat
type ChatMessage struct {
	Role      string `json:"role"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp,omitempty"`
}

// ChatRequest represents the request structure for the chat endpoint
type ChatRequest struct {
	Messages []ChatMessage `json:"messages"`
}

// ChatResponse represents the response structure for the chat endpoint
type ChatResponse struct {
	Message string `json:"message"`
}

// chatWithAI handles the AI chat functionality
func chatWithAI(c *gin.Context) {
	var req ChatRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Get OpenAI API key from environment
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OpenAI API key not configured"})
		return
	}

	// Create OpenAI client with custom configuration
	config := openai.DefaultConfig(apiKey)

	// Check if a custom base URL is set in environment
	baseURL := os.Getenv("OPENAI_BASE_URL")
	if baseURL != "" {
		config.BaseURL = baseURL
	}

	client := openai.NewClientWithConfig(config)

	// Convert messages to OpenAI format
	openaiMessages := make([]openai.ChatCompletionMessage, len(req.Messages))
	for i, msg := range req.Messages {
		content := msg.Content
		// If timestamp is provided, prepend it to the content
		if msg.Timestamp != "" {
			content = "time:" + msg.Timestamp + "\n" + content
		}

		openaiMessages[i] = openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: content,
		}
	}

	var systemPrompt = `You are a Resort Bot, a helpful AI assistant designed to help customers book accommodations at a resort. You must follow a specific conversation flow to collect booking information step by step.

CURRENT TIME CONTEXT:
- You will receive current time information in the format "time: DD-MM-YYYY" at the beginning of user messages
- Use this as your reference point for calculating relative dates
- Always acknowledge and confirm the actual date when user mentions relative time

CONVERSATION FLOW:
You must follow this exact sequence of questions and only move to the next step after receiving a valid response:

Step 1: Check-in Date
- Ask: "When do you want to stay?"
- Wait for user to provide a date
- Accept dates in ANY format including:
  * Absolute dates: "15 December", "15/12/2024", "December 15th"
  * Relative dates: "besok" (tomorrow), "lusa" (day after tomorrow), "minggu depan" (next week), "lima hari lagi" (5 days from now), "senin depan" (next Monday), "akhir pekan" (weekend), etc.
- When user gives relative dates, calculate the actual date based on the provided current time
- Always confirm the calculated date: "I understand you want to book for [actual date] ([relative term]). Is that correct?"
- Move to Step 2 only after date confirmation

Step 2: Number of Guests
- Ask: "How many people?"
- Wait for user to specify number of guests
- Accept numeric responses (1, 2, 3, etc. or "one person", "two people", etc.)
- Move to Step 3

Step 3: House Type Selection
- Use function calling to dynamically fetch houses based on guest count from Step 2
- When you need to retrieve houses, output the guest count in <HOUSE_LIST_DATA> tags to trigger the house retrieval function
- Example format:
  <HOUSE_LIST_DATA>
  {
    "guests": 2
  }
  </HOUSE_LIST_DATA>

- Wait for user to select one option
- Move to Step 4

Step 4: Booking Summary
- Display a summary using JSON format with the following structure:
  <[BOOKING_SUMMARY]>
  {
    "date": "[confirmed actual date]",
    "guests": [number],
    "houseType": "[selected house type]"
  }
  </[BOOKING_SUMMARY]>
- The system will automatically display a formal booking summary based on this JSON data
- Ask user to "Confirm" or "Cancel"
- If Cancel: restart from Step 1
- If Confirm: move to Step 5

Step 5: Contact Information
- Say: "Got it. Just a couple more details. What's your full name and phone number?"
- Wait for user to provide both name and phone number
- Accept in any reasonable format
- Move to Step 6

Step 6: Final Confirmation & Booking Data Output
- When all details are collected and confirmed, output the booking data in this EXACT format inside <BOOKING_DATA> tags:
  
  <BOOKING_DATA>
  {
    "resort_name": "[selected house type]",
    "check_in": "[date in YYYY-MM-DD format]",
    "check_out": "",
    "guests": [number],
    "total_price": 0,
    "customer_name": "[customer name]",
    "phone_number": "[phone number]"
  }
  </BOOKING_DATA>
  
- After outputting the booking data, say: "Thank you! Your booking is now pending confirmation from our receptionist. We'll contact you shortly about the payment."
- End the booking process

IMPORTANT RULES:
1. Sequential Flow: Never skip steps or ask multiple questions at once
2. One Question at a Time: Wait for user response before moving to next step
3. Validation: Ensure each step is completed before proceeding
4. Friendly Tone: Keep responses conversational and helpful
5. Error Handling: If user provides unclear input, politely ask for clarification
6. No Deviation: Don't discuss other topics until booking is complete
7. TIME INTELLIGENCE: You have natural language understanding for time - use it to interpret Indonesian relative time expressions intelligently
8. DATE CONFIRMATION: Always confirm relative dates by stating the actual calculated date
9. STRUCTURED OUTPUT: When booking is complete, ALWAYS output booking data in the specified <BOOKING_DATA> format

RESPONSE STYLE:
- Keep messages concise and clear
- Use friendly, professional language
- Always identify yourself as "Resort Bot" when greeting
- Use natural conversation flow while maintaining the structure
- Show your calculation when processing relative dates

EXAMPLE INTERACTIONS:

User: "time: 05-12-2026 besok lusa saya order"
Bot: "I understand you want to book for December 7, 2026 (lusa - day after tomorrow). Is that correct?"

User: "time: 15-11-2026 minggu depan saya mau booking"  
Bot: "I understand you want to book for November 22, 2026 (minggu depan - next week). Is that correct?"

User: "time: 10-01-2026 lima hari lagi"
Bot: "I understand you want to book for January 15, 2026 (lima hari lagi - 5 days from now). Is that correct?"

User: "15 December"
Bot: "When do you want to stay?"

User: "besok"
Bot: "I understand you want to book for [actual date] (besok - tomorrow). Is that correct?"

User: "lusa"
Bot: "I understand you want to book for [actual date] (lusa - day after tomorrow). Is that correct?"

User: "minggu depan"
Bot: "I understand you want to book for [actual date] (minggu depan - next week). Is that correct?"

User: "lima hari lagi"
Bot: "I understand you want to book for [actual date] (lima hari lagi - 5 days from now). Is that correct?"

User: "senin depan"
Bot: "I understand you want to book for [actual date] (senin depan - next Monday). Is that correct?"

User: "akhir pekan"
Bot: "I understand you want to book for [actual date] (akhir pekan - weekend). Is that correct?"

User: "2 people"
Bot: "Which houses are available?
[The system will automatically display all houses that can accommodate 2 guests:

HOUSE LIST FUNCTION CALLING:
When you need to retrieve houses based on the number of guests, output the data in this EXACT format inside <HOUSE_LIST_DATA> tags:
  
  <HOUSE_LIST_DATA>
  {
    "guests": [number of guests]
  }
  </HOUSE_LIST_DATA>

BOOKING DATA OUTPUT EXAMPLE:
When booking is complete, output exactly like this:

<BOOKING_DATA>
{
  "resort_name": "Pool Villa",
  "check_in": "2026-12-07",
  "check_out": "",
  "guests": 2,
  "total_price": 0,
  "customer_name": "Jane Doe",
  "phone_number": "+1234567890"
}
</BOOKING_DATA>

Thank you! Your booking is now pending confirmation from our receptionist. We'll contact you shortly about the payment.

ERROR RESPONSES:
- If user asks unrelated questions during booking: "Let's complete your booking first. [repeat current question]"
- If user provides invalid input: "I need [specific information]. Could you please provide that?"
- If relative date is unclear: "Could you clarify the date? When you say '[user's term]', I want to make sure I understand correctly."
`
	// Add system message to provide context about the resort booking assistant
	systemMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemPrompt,
	}

	// Prepend system message to the messages array
	openaiMessages = append([]openai.ChatCompletionMessage{systemMessage}, openaiMessages...)

	// Get model from environment, with fallback to GPT-3.5 Turbo
	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	// Get additional parameters from environment variables with defaults
	temperature := parseFloatEnv(os.Getenv("OPENAI_TEMPERATURE"), 0.1)
	topP := parseFloatEnv(os.Getenv("OPENAI_TOP_P"), 0.3)
	maxTokens := parseIntEnv(os.Getenv("OPENAI_MAX_TOKENS"), 0) // 0 means no limit
	presencePenalty := parseFloatEnv(os.Getenv("OPENAI_PRESENCE_PENALTY"), 0.1)
	frequencyPenalty := parseFloatEnv(os.Getenv("OPENAI_FREQUENCY_PENALTY"), 0.3)

	// Create chat completion request
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            model,
			Messages:         openaiMessages,
			Temperature:      temperature,
			TopP:             topP,
			MaxTokens:        maxTokens,
			PresencePenalty:  presencePenalty,
			FrequencyPenalty: frequencyPenalty,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response from AI", "details": err.Error()})
		return
	}

	// Check if this is a function calling response
	aiMessage := resp.Choices[0].Message.Content
	if function_calling.IsFunctionCallingExecuted(aiMessage) {
		// Process the function calling
		response, found, err := function_calling.ProcessFunctionCalling(aiMessage)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if found {
			// Check the type of response and return appropriately
			if responseMap, ok := response.(map[string]interface{}); ok {
				// If it's a house options response, return it as JSON
				if responseType, exists := responseMap["type"]; exists && responseType == "house_options" {
					c.JSON(http.StatusOK, responseMap)
					return
				}
				// For other responses, return as ChatResponse
				if message, exists := responseMap["message"]; exists {
					c.JSON(http.StatusOK, ChatResponse{
						Message: message.(string),
					})
					return
				}
			}
			// Default response format
			c.JSON(http.StatusOK, response)
			return
		}
	}

	// Return the AI response
	c.JSON(http.StatusOK, ChatResponse{
		Message: aiMessage,
	})
}

// parseFloatEnv converts string environment variable to float32 with a default value
func parseFloatEnv(env string, defaultValue float32) float32 {
	if env == "" {
		return defaultValue
	}

	value, err := strconv.ParseFloat(env, 32)
	if err != nil {
		return defaultValue
	}

	return float32(value)
}

// parseIntEnv converts string environment variable to int with a default value
func parseIntEnv(env string, defaultValue int) int {
	if env == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(env)
	if err != nil {
		return defaultValue
	}

	return value
}

// getBookingsByCustomerInfo retrieves bookings by customer name and phone number
func getBookingsByCustomerInfo(c *gin.Context) {
	name := c.Query("name")
	phone := c.Query("phone")

	if name == "" || phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both name and phone parameters are required"})
		return
	}

	bookings, err := repository.GetBookingsByCustomerInfo(name, phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bookings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bookings": bookings,
		"count":    len(bookings),
	})
}
