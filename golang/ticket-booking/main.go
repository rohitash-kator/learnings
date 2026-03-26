package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

const conferenceName = "Go Conference"
const conferenceTickets = 50

var remainingTickets uint = conferenceTickets
var bookings = make([]User, 0)

func main() {
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTicketBooking := validateUserInputs(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketBooking {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v.\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered is too short")
			}
			if !isValidTicketBooking {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total %v of %v tickets are available.\n", remainingTickets, conferenceTickets)
	fmt.Println("Get your ticket here to attend.")
}

func getFirstNames() string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return strings.Join(firstNames, ", ")
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	// Create a map for a user
	user := User{
		numberOfTickets: userTickets,
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
	}

	bookings = append(bookings, user)

	fmt.Printf("List of bookings is %v.\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("--------------------------------------------------------")
	fmt.Printf("Sending ticket:\n%v to email address %v\n", ticket, email)
	fmt.Println("--------------------------------------------------------")
	wg.Done()
}
