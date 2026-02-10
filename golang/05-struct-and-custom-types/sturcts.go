package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Factory/Constructor/Creation function to create a new user struct instance (Note: This is a regular function, not a method)
// Important: This function returns a pointer to the user struct instance so that the changes are reflected in the original struct instance,
// otherwise the changes will be reflected in a copy of the struct instance and not the original struct instance
func newUser(firstName string, lastName string, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
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
	appUser = newUser(userFirstName, userLastName, userBirthDate)

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

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string

	fmt.Scan(&value)

	return value
}
