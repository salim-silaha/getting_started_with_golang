package main

import (
	"fmt"
	"strings"
)

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

	// Introduction of for loop
	for {
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

		// using if statement to validate user input.
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			// logic for calculating remaining tickets
			remainingTickets = remainingTickets - userTickets

			// logic for booking the tickets.
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank You %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// introduction of an array to hold the values of the firstnames.
			firstNames := []string{}

			// for loop to loop through the bookings array.
			// more or less like a for each loop.
			// it get the value of the first name from the bookings array and appends/adds the firstnames to
			// a different array that holds only firstnames.
			// _ is a technique to tell the for loop to ignore the first argurment.
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			// printing out all the bookings done.
			fmt.Printf("This are the first names of all our bookings: %v\n", firstNames)
		} else {
			if !isValidName {
				fmt.Println("Name provided is invalid, please check the password and try again.")
			}
			if !isValidEmail {
				fmt.Println("The email provided is not valid.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
			continue
		}
	}
}
