package main

import (
	"booking-app/model"
	_ "booking-app/model"
	"fmt"
)

var conferenceName = "GO CONFERENCE"

// User info
var bookingSystemCapacity uint
var reservations []model.Reservation

func main() {

	fmt.Println("==========================", conferenceName, "==========================")
	fmt.Printf("Welcome to our booking application!\n")

	bookingSystemCapacity = defineSystemConfig()

	fmt.Printf("\nGet your ticket here to attend\n")

	fmt.Println("==========================")

	for {
		if bookingSystemCapacity == 0 {
			fmt.Println("There's no tickets available")
			break
		}
		var user = obtainUserInfo()
		var reservation = obtainBooking(user)

		reservations = append(reservations, reservation)
		bookingSystemCapacity = updateRemainingTickets(bookingSystemCapacity, reservation.Amount)
		printBooking(reservation)
	}

	fmt.Println("==========================")

	fmt.Printf("\n\nATTENTION: Are still %v tickets available.\n\n", bookingSystemCapacity)
	fmt.Println("==========================")
}

func obtainBooking(user model.User) model.Reservation {
	var reservation model.Reservation
	reservation.User = user
	var amount uint
	fmt.Print("How many tickets to book?: ")
	_, err := fmt.Scan(&amount)
	if err != nil {
		return reservation
	}
	if amount > bookingSystemCapacity {
		panic("Invalid amount. Please try again...")
	}
	reservation.Amount = amount
	return reservation
}

func defineSystemConfig() uint {
	fmt.Print("Define the number of tickets will be available: ")
	var systemCapacity uint
	_, err := fmt.Scan(&systemCapacity)
	if err != nil {
		return systemCapacity
	}

	return systemCapacity
}

func obtainUserInfo() model.User {

	var user model.User

	fmt.Print("Enter your first name: ")
	_, err := fmt.Scan(&user.FirstName)
	if err != nil {
		return user
	}
	fmt.Print("Enter your last name: ")
	_, err = fmt.Scan(&user.LastName)
	if err != nil {
		return user
	}
	fmt.Print("Enter your email address: ")
	_, err = fmt.Scan(&user.Email)
	if err != nil {
		return user
	}

	return user
}

func updateRemainingTickets(capacity uint, amount uint) uint {
	return capacity - amount
}

func printBooking(reservation model.Reservation) {
	fmt.Printf("\n\"Thank you %v %v for booking %v tickets. You will receive an email confirmation at %v\" \n", reservation.User.FirstName, reservation.User.LastName, reservation.Amount, reservation.User.Email)
}
