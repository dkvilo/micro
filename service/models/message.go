package models

// Message model
// swagger:response Message
type Message struct {
	// Sample message from server
	// 
	// Required: true
	// example: Hello, World
	Message string	`json:"message"`
	// Port number of the server 
	// 
	// Required: true
	// example: :8080
	Port 		string	`json:"port"`
}

