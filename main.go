package main

// introduction of go packages to wrap up packages and calling them in the other packages.
// for the package getting_started_with_golang/header
// we import the header package and a function called validateUserInput.
// on the main function we invoke the function using the "package_name.Function" which we exposed
// in the other package.
import (
	"fmt"
	"getting_started_with_golang/header"
	"strings"
)

// Introduction of variables.
// Conference has a data type of int.
const conferenceTickets int = 50

// constant variable of the name of the conference.
var conferenceName = "Go Conference"

// variable to hold the value of the remaining tickets.
var remainingTickets uint = 50

// Introduction of array
var bookings = []string{}

func main() {

	greetUsers()

	// Introduction of for loop
	for {

		// declaration of the returned values from the functions.
		firstName, lastName, email, userTickets := getUserInput()

		// declaration and instantiation of the returned values from the functions.
		isValidName, isValidEmail, isValidTicketNumber := header.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// instatiation of the bookTickets function with passing the argurments.
			bookTickets(firstName, lastName, userTickets, email)

			firstNames := printFirstNames()

			// printing out all the bookings done.
			fmt.Printf("This are the first names of all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the program
				println("All tickets have been booked today. Wait again next time")
				break
			}

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

// declaration of functions.

// the argurments conferenceName, conferenceTickets and remainingTickets are now accessible inside the function because of the global declaration.
func greetUsers() {
	// Printing the variables and a welcome message.
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func getUserInput() (string, string, string, uint) {

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

	// returning this value to the caller.
	// go is the only programming language that allows multiple return.
	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, userTickets uint, email string) {

	// logic for calculating remaining tickets
	remainingTickets = remainingTickets - userTickets

	// logic for booking the tickets.
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank You %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func printFirstNames() []string {
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
	return firstNames
}
