package main

import (
	"encoding/json"
	"log"
	"net/http"
	"verifier/utils"
)

// VerifyRequest represents the structure of the incoming JSON request
type VerifyRequest struct {
	Email string `json:"email"`
}

// VerifyResponse represents the structure of the JSON response
type VerifyResponse struct {
	Email   string `json:"email"`
	Valid   bool   `json:"valid"`
	Message string `json:"message,omitempty"`
}

// verifyEmailHandler handles the verification of email addresses
func verifyEmailHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	var request VerifyRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if request.Email == "" {
		http.Error(w, "Email address cannot be empty", http.StatusBadRequest)
		return
	}

	valid, _ := utils.IsValidEmail(request.Email)

	response := VerifyResponse{
		Email: request.Email,
		Valid: valid,
	}

	if !valid {
		response.Message = "Invalid email address or cannot be verified!"
	}

	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/verify-email", verifyEmailHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
