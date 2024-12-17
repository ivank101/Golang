package main

import (
	"fmt"
	"strings"
)

const concertTickets int = 50

var concertName = "Holo Concert"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserinput(firstName, lastName, email, userTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first name of our bookings are : %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our concert is booked out")
				break
			}
		} else {
			if !isValidEmail {
				fmt.Println("The name that you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email that you entered is not correct")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", concertName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", concertTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, concertName)

}
