package main

import "fmt"

func main() {

	// Introduction of variables.
	//Conference has a data type of int.
	const conferenceTickets int = 50

	// constant variable of the name of the conference.
	conferenceName := "Go Conference"

	// variable to hold the value of the remaining tickets.
	var remainingTickets uint = 50

	// Introduction of array
	bookings := []string{}

	// Printing the variables and a welcome message.
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	// This are variables whose values are going to be filled by the user.
	var firstName string
	var lastName string
	var email string
	// userTicket is not int because it cannot be decimal, int always for decimal.
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	// getting the values supplied by the user.
	// &firstName is a pointer to the memory location of the variable.
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	// getting the values supplied by the user.
	// &lastName is a pointer to the memory location of the variable.
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email Address: ")
	// getting the values supplied by the user.
	// &email is a pointer to the memory location of the variable.
	fmt.Scanln(&email)

	fmt.Println("Enter Your Number of tickets you want: ")
	// getting the values supplied by the user.
	// &userTickets is a pointer to the memory location of the variable.
	fmt.Scanln(&userTickets)

	// logic for calculating remaining tickets
	remainingTickets = remainingTickets - userTickets

	// logic for booking the tickets.
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank You %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// printing out all the bookings done.
	fmt.Printf("This are all our bookings: %v\n", bookings)
}
