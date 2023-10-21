package main

import (
	"booking-app/model"
	_ "booking-app/model"
	"fmt"
)

var conferenceName = "GO CONFERENCE"

// User info
var numberOfTickets uint
var remainingTickets uint = 50

var user model.User

func main() {

	fmt.Println("==========================", conferenceName, "==========================")
	fmt.Printf("Welcome to our booking application!\n")
	fmt.Printf("\nGet your ticket here to attend\n")

	// Ask user their name

	fmt.Println("==========================")

	user := obtainUserInfo(user)

	fmt.Println("==========================")

	obtainConferenceData()

	updateRemainingTickets(numberOfTickets)

	fmt.Printf("\n\"Thank you %v %v for booking %v tickets. You will receive an email confirmation at %v\" \n", user.FirstName, user.LastName, numberOfTickets, user.Email)

	fmt.Printf("\n\nATTENTION: Are still %v tickets available.\n\n", remainingTickets)
	fmt.Println("==========================")
}

func obtainUserInfo(u model.User) model.User {

	fmt.Print("Enter your first name: ")
	_, err := fmt.Scan(&u.FirstName)
	if err != nil {
		return u
	}
	fmt.Print("Enter your last name: ")
	_, err = fmt.Scan(&u.LastName)
	if err != nil {
		return u
	}
	fmt.Print("Enter your email address: ")
	_, err = fmt.Scan(&u.Email)
	if err != nil {
		return u
	}

	return u

}

func obtainConferenceData() {
	fmt.Print("How many tickets to book?: ")
	fmt.Scan(&numberOfTickets)
}

func updateRemainingTickets(amount uint) {
	remainingTickets = remainingTickets - amount
}
