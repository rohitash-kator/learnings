package dao

// User Model
type User struct {
	ID   int
	Name string
}

// Interface Definition
type UserDaoInterface interface {
	GetUser(ID int) (User, error)
}
