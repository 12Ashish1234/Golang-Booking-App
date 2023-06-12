package main

import (
	"strings"
)

func CheckValidity(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	// In GO you can return more than one value to the calling function
	return isValidName, isValidEmail, isValidUserTickets
}
