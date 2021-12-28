package main

import (
	"fmt"
	"strings"
	"booking-app/helper"
)

var conferenceName = "Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = []string{}


func main() {

	greetUsers()

	for {
		firstName, lastName , email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber :=  helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Println(firstNames)
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out, Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered us too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered udoesn't contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("We have total of %v tickets %v tickets are remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}


func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}



func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets (userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("user %v booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}