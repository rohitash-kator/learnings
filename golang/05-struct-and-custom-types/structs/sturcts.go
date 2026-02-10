package main

import (
	"fmt"

	"example.com/go-structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

	var appUser *user.User

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
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		// panic(err)
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "admin123")

	// admin.User.OutputUserDetails()
	// admin.User.ClearUserName()
	// admin.User.OutputUserDetails()

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// Direct assignment of values to the struct fields using variable names and values (order is important)
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthDate,
	// 	time.Now(),
	// }

	// outputUserDetails(&appUser)
	appUser.OutputUserDetails() // Call the method on the struct instance
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string

	fmt.Scanln(&value)

	return value
}
