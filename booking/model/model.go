package model

type User struct {
	FirstName string
	LastName  string
	Email     string
}

type Reservation struct {
	User   User
	Amount uint
}
