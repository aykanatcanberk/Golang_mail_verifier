package utils

import (
	"net"
	"regexp"
	"strings"
)

func IsValidEmail(email string) (bool, error) {

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false, err 
		return false, nil 
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false, nil 
	}

	// Assign username and domain parts to separate variables.
	username := parts[0]
	domain := parts[1]

	// Check the length of username.
	if len(username) < 1 || len(username) > 64 {
		return false, nil 
	}

	// Check the length of domain.
	if len(domain) < 3 || len(domain) > 255 {
		return false, nil 
	}

	// Check for valid MX (Mail Exchange) records for the domain.
	_, mxErr := net.LookupMX(domain)
	if mxErr != nil {
		return false, mxErr 
	}

	// Check for valid NS (Name Server) records for the domain.
	_, nsErr := net.LookupNS(domain)
	if nsErr != nil {
		return false, nsErr 
	}
	return true, nil
}
