package main

import (
	"fmt"
	"verifier/utils"
)

func main() {
	// Define email addresses to be tested
	emails := []string{
		"test@gmail.com",
		"invalid.email.com",
		"missingatexample.com",
	}

	// Call IsValidEmail function for each email address and check the result
	for _, email := range emails {
		valid, err := utils.IsValidEmail(email)
		if err != nil {
			fmt.Printf("An error occurred while verifying the email '%s': %v\n", email, err)
			continue
		}
		if valid {
			fmt.Printf("The email address '%s' appears to be valid and existent.\n", email)
		} else {
			fmt.Printf("The email address '%s' is invalid or may not exist.\n", email)
		}
	}
}
