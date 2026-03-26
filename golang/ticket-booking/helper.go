package main

import "strings"

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketBooking := userTickets > 0 && remainingTickets >= userTickets

	return isValidName, isValidEmail, isValidTicketBooking
}
