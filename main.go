package main

import "fmt"

func main() {

	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets, and", remainingTickets, "are still available!")
	fmt.Println("Get your ticket here to attend")

	var userName string
	var userTicket int
	// Ask user their name

	userName = "Lopes"
	userTicket = 10
	fmt.Printf("The user %v booked the %v° ticket\n", userName, userTicket)
}
