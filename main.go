// main package of program
// all code must be in a package
package main

import "fmt"

// to let go know where to start executing
func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("conference tickets is %T", conferenceTickets)
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")


	// var bookings = [50]string{"Nicole", "David", "James"}
	var bookings []string
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)
	
	fmt.Printf("Thank you %v for booking %v tickets\n", bookings[0], userTickets)
	fmt.Printf("Remaining Tickets: %v\n", remainingTickets)

	fmt.Printf("These are all the booking %v", bookings)
}
