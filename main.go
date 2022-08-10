// main package of program
// all code must be in a package
package main

import (
	"fmt"
	"strings"
)

// to let go know where to start executing
func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conference tickets is %T", conferenceTickets)
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

	for {
		// var bookings = [50]string{"Nicole", "David", "James"}
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
		
		if userTickets < remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)
			
			fmt.Printf("Thank you %v for booking %v tickets\n", bookings[0], userTickets)
			fmt.Printf("Remaining Tickets: %v\n", remainingTickets)
	
			firstNames := []string{}
			for _, booking := range  bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all the booking on first names basis %v\n", firstNames)
	
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end the program
				fmt.Printf("Conference is booked out")
				break
			}	
		} else if userTickets == remainingTickets {
			// do something else

		} else {
			fmt.Printf("You cannot book more that the available tickets which is %v\n", remainingTickets)
			continue
		}
	}
	
}
