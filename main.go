package main

import "fmt"

var conferenceName = "GO CONFERENCE"

// User info
var firstName string
var lastName string
var email string
var numberOfTickets uint

var remainingTickets uint = 50

func main() {

	fmt.Println("==========================", conferenceName, "==========================")
	fmt.Printf("Welcome to our booking application!\n")
	fmt.Printf("\nGet your ticket here to attend\n")

	// Ask user their name

	fmt.Println("==========================")

	obtainUserInfo()

	fmt.Println("==========================")

	obtainConferenceData()

	updateRemainingTickets(numberOfTickets)

	fmt.Printf("\n\"Thank you %v %v for booking %v tickets. You will receive an email confirmation at %v\" \n", firstName, lastName, numberOfTickets, email)

	fmt.Printf("\n\nATTENTION: Are still %v tickets available.\n\n", remainingTickets)
	fmt.Println("==========================")
}

func obtainUserInfo() {

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

}

func obtainConferenceData() {
	fmt.Print("How many tickets to book?: ")
	fmt.Scan(&numberOfTickets)
}

func updateRemainingTickets(amount uint) {
	remainingTickets = remainingTickets - amount
}
