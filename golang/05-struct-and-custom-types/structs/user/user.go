package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	// User     User // Embedding the User struct as a field
	User // Embedding the User struct as a field anonymously (Note: This is the same as the above commented out code)
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
func (userInfo *User) OutputUserDetails() {
	// ...
	// Dereference the pointer to get the value of the struct fields using the dot operator (dereferencing is optional)
	fmt.Println(userInfo.firstName, userInfo.lastName, userInfo.birthDate)
}

// Attaching a mutating method to a struct
// Important: This method takes a pointer to the struct as the receiver so that the changes are reflected in the original struct instance,
// otherwise the changes will be reflected in a copy of the struct instance and not the original struct instance
func (userInfo *User) ClearUserName() {
	userInfo.firstName = ""
	userInfo.lastName = ""

}

func NewAdmin(email, password string) Admin {

	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}
}

// Factory/Constructor/Creation function to create a new user struct instance (Note: This is a regular function, not a method)
// Important: This function returns a pointer to the user struct instance so that the changes are reflected in the original struct instance,
// otherwise the changes will be reflected in a copy of the struct instance and not the original struct instance
func New(firstName string, lastName string, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, Last name and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
