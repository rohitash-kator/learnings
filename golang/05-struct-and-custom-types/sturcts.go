package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

	var appUser *user

	// Initialize the struct with no values (all fields will be set to their zero/nil values)
	// appUser = user{}

	// Initialize the struct with values for some fields and the rest will be set to their zero/nil values
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	// birthDate: userBirthDate,
	// 	createdAt: time.Now(),
	// }

	// Direct assignment of values to the struct fields using variable names and values (order is not important)
	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		// panic(err)
		fmt.Println(err)
		return
	}

	// Direct assignment of values to the struct fields using variable names and values (order is important)
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthDate,
	// 	time.Now(),
	// }

	// outputUserDetails(&appUser)
	appUser.outputUserDetails() // Call the method on the struct instance

	appUser.clearUserName()

	appUser.outputUserDetails()
}

// Using value receiver
// func outputUserDetails(userInfo user) {
// 	// ...
// 	fmt.Println(userInfo.firstName, userInfo.lastName, userInfo.birthDate)
// }

// Using pointer receiver
// func outputUserDetails(userInfo *user) {
// 	// ...
// 	// Dereference the pointer to get the value of the struct fields using the dot operator (dereferencing is optional)
// 	fmt.Println((*userInfo).firstName, userInfo.lastName, userInfo.birthDate)
// }

// Attaching a method to a struct
func (userInfo user) outputUserDetails() {
	// ...
	// Dereference the pointer to get the value of the struct fields using the dot operator (dereferencing is optional)
	fmt.Println(userInfo.firstName, userInfo.lastName, userInfo.birthDate)
}

// Attaching a mutating method to a struct
// Important: This method takes a pointer to the struct as the receiver so that the changes are reflected in the original struct instance,
// otherwise the changes will be reflected in a copy of the struct instance and not the original struct instance
func (userInfo *user) clearUserName() {
	userInfo.firstName = ""
	userInfo.lastName = ""

}

// Factory/Constructor/Creation function to create a new user struct instance (Note: This is a regular function, not a method)
// Important: This function returns a pointer to the user struct instance so that the changes are reflected in the original struct instance,
// otherwise the changes will be reflected in a copy of the struct instance and not the original struct instance
func newUser(firstName string, lastName string, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, Last name and birth date are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string

	fmt.Scanln(&value)

	return value
}
