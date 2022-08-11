// main package of program
// all code must be in a package
package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// package level variable (cannot use :=)
var conferenceName string = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0) // var bookings = [50]string{"Nicole", "David", "James"}

type UserData struct {
	firstName       string
	lastName        string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

// to let go know where to start executing
func main() {
	greetUsers()
	// by default - for {} == for true {}
	// for {
	firstName, lastName, email, userTickets := getUserInputs()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName)

		firstNames := getFirstNames(bookings)
		fmt.Printf("These are all the booking on first names basis %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			// end the program
			fmt.Printf("Conference is booked out")
			// break
		}
	} else {
		fmt.Println("Your input data is invalid, try again")
		// continue
	}
	wg.Wait()
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}

func bookTicket(userTickets uint, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets

	// create Map
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v for booking %v tickets\n", bookings[0], userTickets)
	fmt.Printf("Remaining Tickets: %v\n", remainingTickets)
}

func getFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func sendTicket(userTickets uint, firstName string, lastName string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("sending ticket:\n%v to email\n", ticket)
	fmt.Println("##############")
	wg.Done()
}

/**
city := "London"

switch city {
	case "New York":
		// execute code for new york
	case "Singapore", "Hong Kong":
		// execute some code

	default:
		// fmt.Println("no valid city selected")
}
**/
