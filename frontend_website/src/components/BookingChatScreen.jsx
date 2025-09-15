import React, { useState } from 'react';

const BookingChatScreen = ({ navigateTo }) => {
  // Use a counter for generating unique IDs
  const [messageIdCounter, setMessageIdCounter] = useState(2);
  
  // Function to format date in a more readable way for AI
  const formatReadableDate = (date) => {
    return date.toLocaleDateString('en-US', { 
      weekday: 'long', 
      year: 'numeric', 
      month: 'long', 
      day: 'numeric' 
    });
  };

  const formatTime = (date) => {
    return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
  };

  const formatDateTime = (date) => {
    return date.toLocaleDateString('en-US', { 
      weekday: 'short', 
      month: 'short', 
      day: 'numeric',
      hour: '2-digit', 
      minute: '2-digit' 
    });
  };

  const [messages, setMessages] = useState([
    {
      id: 1,
      sender: 'bot',
      text: `Hello! I am your resort booking assistant. When do you want to stay?`,
      avatar: (
        <svg className="h-8 w-8 text-[#38e07b]" fill="none" height="32" viewBox="0 0 24 24" width="32" xmlns="http://www.w3.org/2000/svg">
          <path d="M12 2L1 9L4 22H20L23 9L12 2Z" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
          <path d="M12 2V22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
          <path d="M23 9H1" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
          <path d="M4 22L12 15L20 22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
          <path d="M10 12H14" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
        </svg>
      ),
      timestamp: new Date()
    }
  ]);

  const [inputValue, setInputValue] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const handleSendMessage = async () => {
    if (inputValue.trim() === '' || isLoading) return;
    
    // Generate unique IDs for messages
    const newUserMessageId = messageIdCounter;
    
    // Add user message to chat
    const userMessage = {
      id: newUserMessageId,
      sender: 'user',
      text: inputValue,
      timestamp: new Date()
    };
    
    const newMessages = [...messages, userMessage];
    setMessages(newMessages);
    setInputValue('');
    setIsLoading(true);
    setMessageIdCounter(messageIdCounter + 1); // Update counter for next messages
    
    try {
      // Prepare messages for AI (excluding UI-specific properties)
      const aiMessages = newMessages
        .filter(msg => msg.sender === 'user' || msg.sender === 'bot')
        .map(msg => ({
          role: msg.sender === 'user' ? 'user' : 'assistant',
          content: msg.text,
          // Include timestamp for backend processing
          ...(msg.sender === 'user' && {
            timestamp: new Date().toLocaleDateString('en-US', { 
              weekday: 'long', 
              year: 'numeric', 
              month: 'long', 
              day: 'numeric',
              hour: '2-digit',
              minute: '2-digit',
              hour12: true
            })
          })
        }));
      
      // Use the current domain for API calls
      const apiUrl = `${window.location.origin}`;
      
      // Send to backend AI service
      const response = await fetch(`${apiUrl}/api/chat/message`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ messages: aiMessages }),
      });
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const data = await response.json();
      
      // Generate a new unique ID for the AI message
      const newAiMessageId = messageIdCounter + 1;
      
      // Handle different response types
      if (data.type === 'house_options') {
        // Parse the houses JSON string
        const houses = JSON.parse(data.houses);
        
        // Add AI response with house options to chat
        const aiMessage = {
          id: newAiMessageId,
          sender: 'bot',
          text: data.message,
          houses: houses, // Include houses data
          timestamp: new Date()
        };
        
        setMessages(prevMessages => [...prevMessages, aiMessage]);
      } else {
        // Check if the message contains booking summary data
        const bookingSummary = extractBookingSummary(data.message);
        
        if (bookingSummary) {
          // Add the booking summary as a regular message in the chat, but hide the JSON part
          const cleanMessage = data.message.replace(/<\[BOOKING_SUMMARY\]>.*?<\/\[BOOKING_SUMMARY\]>/s, '').trim();
          
          const aiMessage = {
            id: newAiMessageId,
            sender: 'bot',
            text: cleanMessage,
            bookingSummary: bookingSummary,
            avatar: (
              <svg className="h-8 w-8 text-[#38e07b]" fill="none" height="32" viewBox="0 0 24 24" width="32" xmlns="http://www.w3.org/2000/svg">
                <path d="M12 2L1 9L4 22H20L23 9L12 2Z" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                <path d="M12 2V22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                <path d="M23 9H1" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                <path d="M4 22L12 15L20 22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                <path d="M10 12H14" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
              </svg>
            ),
            timestamp: new Date()
          };
          
          setMessages(prevMessages => [...prevMessages, aiMessage]);
        } else {
          // Add regular AI response to chat
          const aiMessage = {
            id: newAiMessageId,
            sender: 'bot',
            text: data.message,
            avatar: (
              <svg className="h-8 w-8 text-[#38e07b]" fill="none" height="32" viewBox="0 0 24 24" width="32" xmlns="http://www.w3.org/2000/svg">
                <path d="M12 2L1 9L4 22H20L23 9L12 2Z" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                <path d="M12 2V22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                <path d="M23 9H1" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                <path d="M4 22L12 15L20 22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                <path d="M10 12H14" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
              </svg>
            ),
            timestamp: new Date()
          };
          
          setMessages(prevMessages => [...prevMessages, aiMessage]);
        }
      }
      
      setMessageIdCounter(newAiMessageId + 1); // Update counter for next messages
    } catch (error) {
      console.error('Error sending message to AI:', error);
      
      // Generate a new unique ID for the error message
      const newErrorMessageId = messageIdCounter + 1;
      
      // Add error message to chat
      const errorMessage = {
        id: newErrorMessageId,
        sender: 'bot',
        text: 'Sorry, I encountered an error. Please try again.',
        avatar: (
          <svg className="h-8 w-8 text-[#38e07b]" fill="none" height="32" viewBox="0 0 24 24" width="32" xmlns="http://www.w3.org/2000/svg">
            <path d="M12 2L1 9L4 22H20L23 9L12 2Z" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
            <path d="M12 2V22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
            <path d="M23 9H1" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
            <path d="M4 22L12 15L20 22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
            <path d="M10 12H14" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
          </svg>
        ),
        timestamp: new Date()
      };
      
      setMessages(prevMessages => [...prevMessages, errorMessage]);
      setMessageIdCounter(newErrorMessageId + 1); // Update counter for next messages
    } finally {
      setIsLoading(false);
    }
  };

  const handleKeyPress = (e) => {
    if (e.key === 'Enter') {
      handleSendMessage();
    }
  };

  // Function to handle house selection
  const handleHouseSelection = (house) => {
    // Send the selected house name as a message
    const selectionMessage = `I choose the ${house.name}`;
    setInputValue(selectionMessage);
    // Small delay to ensure state is updated before sending
    setTimeout(() => {
      handleSendMessage();
    }, 100);
  };

  // Function to handle booking summary confirmation
  const handleBookingConfirm = () => {
    // Send confirmation message
    const confirmMessage = "Confirm";
    setInputValue(confirmMessage);
    setTimeout(() => {
      handleSendMessage();
    }, 100);
  };

  // Function to handle booking summary cancellation
  const handleBookingCancel = () => {
    // Send cancellation message
    const cancelMessage = "Cancel";
    setInputValue(cancelMessage);
    setTimeout(() => {
      handleSendMessage();
    }, 100);
  };

  // Function to extract booking summary data from message
  const extractBookingSummary = (message) => {
    const startTag = "<[BOOKING_SUMMARY]>";
    const endTag = "</[BOOKING_SUMMARY]>";
    
    const startIndex = message.indexOf(startTag);
    if (startIndex === -1) return null;
    
    const endIndex = message.indexOf(endTag);
    if (endIndex === -1) return null;
    
    const jsonStart = startIndex + startTag.length;
    const jsonEnd = endIndex;
    const jsonString = message.substring(jsonStart, jsonEnd).trim();
    
    try {
      return JSON.parse(jsonString);
    } catch (e) {
      console.error("Error parsing booking summary JSON:", e);
      return null;
    }
  };

  return (
    <div className="relative flex size-full min-h-screen flex-col justify-between overflow-x-hidden bg-white">
      <div className="flex-1 overflow-y-auto">
        <header className="sticky top-0 z-10 flex items-center bg-white/80 backdrop-blur-sm p-4 pb-3 justify-between border-b border-gray-200">
          <button 
            className="text-gray-700 flex size-10 items-center justify-center"
            onClick={() => navigateTo && navigateTo('home')}
          >
            <svg fill="currentColor" height="24" viewBox="0 0 256 256" width="24" xmlns="http://www.w3.org/2000/svg">
              <path d="M224,128a8,8,0,0,1-8,8H59.31l58.35,58.34a8,8,0,0,1-11.32,11.32l-72-72a8,8,0,0,1,0-11.32l72-72a8,8,0,0,1,11.32,11.32L59.31,120H216A8,8,0,0,1,224,128Z"></path>
            </svg>
          </button>
          <h1 className="text-gray-900 text-lg font-bold leading-tight tracking-[-0.015em]">Booking Assistant</h1>
          <div className="w-10"></div>
        </header>
        <main className="p-4 space-y-6">
          {messages.map((message) => (
            <div key={message.id} className={`flex ${message.sender === 'user' ? 'justify-end' : 'items-end gap-2.5 flex items-start'}`}>
              {message.sender !== 'user' && (
                <div className="size-8 rounded-full">
                  <svg className="h-8 w-8 text-[#38e07b]" fill="none" height="32" viewBox="0 0 24 24" width="32" xmlns="http://www.w3.org/2000/svg">
                    <path d="M12 2L1 9L4 22H20L23 9L12 2Z" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                    <path d="M12 2V22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                    <path d="M23 9H1" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                    <path d="M4 22L12 15L20 22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                    <path d="M10 12H14" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                  </svg>
                </div>
              )}
              <div className={`flex flex-col max-w-[85%] gap-1 ${message.sender === 'user' ? 'items-end' : 'items-start'}`}>
                {message.sender !== 'user' && (
                  <p className="text-gray-600 text-xs font-medium">Resort Assistant</p>
                )}
                <div className={`rounded-2xl px-4 py-3 max-w-[100%] ${message.sender === 'user' ? 'bg-green-500 text-white rounded-br-lg' : 'bg-gray-100 text-gray-900 rounded-bl-lg'}`}>
                  {message.text}
                  {/* Render house options if they exist */}
                  {message.houses && message.houses.length > 0 && (
                    <div className="mt-3 space-y-3">
                      {message.houses.map((house) => (
                        <div
                          key={house.id}
                          className="border border-gray-200 rounded-lg p-3 bg-white cursor-pointer hover:bg-gray-50 transition-colors"
                          onClick={() => handleHouseSelection(house)}
                        >
                          <div className="flex items-center gap-3">
                            <img
                              src={house.image_url}
                              alt={house.name}
                              className="w-16 h-16 rounded-md object-cover"
                            />
                            <div className="flex-1">
                              <h3 className="font-semibold text-gray-900">{house.name}</h3>
                              <p className="text-sm text-gray-600">{house.guests} guests • ${house.price_per_night}/night</p>
                            </div>
                          </div>
                        </div>
                      ))}
                    </div>
                  )}
                  
                  {/* Render booking summary if it exists */}
                  {message.bookingSummary && (
                    <div className="mt-4 border border-gray-200 rounded-lg p-4 bg-white">
                      <h3 className="text-lg font-semibold text-gray-800 mb-3">Booking Summary</h3>
                      <div className="space-y-2 mb-4">
                        <div className="flex justify-between">
                          <span className="text-gray-600">Date:</span>
                          <span className="font-medium">{message.bookingSummary.date}</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-gray-600">Guests:</span>
                          <span className="font-medium">{message.bookingSummary.guests} people</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-gray-600">House Type:</span>
                          <span className="font-medium">{message.bookingSummary.houseType}</span>
                        </div>
                      </div>
                      <div className="flex space-x-3">
                        <button
                          onClick={handleBookingCancel}
                          className="flex-1 bg-gray-200 hover:bg-gray-300 text-gray-800 font-medium py-2 px-3 rounded-lg transition duration-200 text-sm"
                        >
                          Cancel
                        </button>
                        <button
                          onClick={handleBookingConfirm}
                          className="flex-1 bg-green-500 hover:bg-green-600 text-white font-medium py-2 px-3 rounded-lg transition duration-200 text-sm"
                        >
                          Confirm
                        </button>
                      </div>
                    </div>
                  )}
                </div>
                <p className="text-gray-500 text-xs mt-1">
                  {formatDateTime(message.timestamp)}
                </p>
              </div>
            </div>
          ))}
          {isLoading && (
            <div className="flex items-end gap-2.5 flex items-start">
              <div className="size-8 rounded-full">
                <svg className="h-8 w-8 text-[#38e07b]" fill="none" height="32" viewBox="0 0 24 24" width="32" xmlns="http://www.w3.org/2000/svg">
                  <path d="M12 2L1 9L4 22H20L23 9L12 2Z" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                  <path d="M12 2V22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                  <path d="M23 9H1" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                  <path d="M4 22L12 15L20 22" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                  <path d="M10 12H14" stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5"></path>
                </svg>
              </div>
              <div className="flex flex-col max-w-[85%] gap-1 items-start">
                <p className="text-gray-600 text-xs font-medium">Resort Assistant</p>
                <div className="rounded-2xl px-4 py-3 max-w-[100%] bg-gray-100 text-gray-900 rounded-bl-lg">
                  <div className="flex space-x-2">
                    <div className="w-2 h-2 rounded-full bg-gray-400 animate-bounce"></div>
                    <div className="w-2 h-2 rounded-full bg-gray-400 animate-bounce" style={{ animationDelay: '0.2s' }}></div>
                    <div className="w-2 h-2 rounded-full bg-gray-400 animate-bounce" style={{ animationDelay: '0.4s' }}></div>
                  </div>
                </div>
                <p className="text-gray-500 text-xs mt-1">
                  {formatDateTime(new Date())}
                </p>
              </div>
            </div>
          )}
        </main>
      </div>
      <footer className="sticky bottom-0 bg-white/80 backdrop-blur-sm pt-3 pb-4 px-4 border-t border-gray-200">
        <div className="flex items-center gap-3">
          <input 
            className="flex-1 bg-gray-100 text-gray-900 placeholder:text-gray-500 rounded-full h-12 px-5 focus:outline-none focus:ring-2 focus:ring-green-500" 
            placeholder="Type a message..." 
            type="text"
            value={inputValue}
            onChange={(e) => setInputValue(e.target.value)}
            onKeyPress={handleKeyPress}
            disabled={isLoading}
          />
          <button 
            className={`bg-green-500 text-white size-12 rounded-full flex items-center justify-center shrink-0 hover:bg-green-600 ${isLoading ? 'opacity-50 cursor-not-allowed' : ''}`}
            onClick={handleSendMessage}
            disabled={isLoading}
          >
            {isLoading ? (
              <svg className="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
            ) : (
              <svg fill="currentColor" height="24" viewBox="0 0 256 256" width="24" xmlns="http://www.w3.org/2000/svg">
                <path d="M221.66,133.66l-72,72a8,8,0,0,1-11.32-11.32L196.69,136H40a8,8,0,0,1,0-16H196.69L138.34,61.66a8,8,0 0,1,11.32-11.32l72,72A8,8,0,0,1,221.66,133.66Z"></path>
              </svg>
            )}
          </button>
        </div>
      </footer>
    </div>
  );
};

export default BookingChatScreen;