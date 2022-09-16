package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var emailAddress string
		var userTickets uint

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address")
		fmt.Scan(&emailAddress)
		fmt.Println("Enter number of tickets you're buying:)")
		fmt.Scan(&userTickets)

		//fmt.Scan introduced to coolect & store introduced variables = firstName, lastName

		remainingTickets = remainingTickets - userTickets

		//We get a list of "firstName" & "lastName" from each "bookings". var bookings comes from firstName & Name. append describes = list
		//ARRAYS are fixed(50), slices are open-ended since the list can have tens of thousands of people.
		//Slice= Nicole Smith. Elements= Nicole, Smith
		//Slices are data types that are better to use compared to arrays. Arrays are fixed, slices open-ended, good for lists.

		fmt.Printf("Hey! %v %v you've successfully booked %v tickets You'll receive a confirmation message at %v with your ticket details:).\n", firstName, lastName, userTickets, emailAddress)
		fmt.Printf("We've got %v tickets remaining for the %v !\n", remainingTickets, conferenceName)

		//[]string{} is used because the variable will be filled, we currently don't know it, means the list is empty - no pre-defined variables:
		//index numbering the list 0-10+ ; booking = Nicole Smith; because bookings = "Nicole Smith Nana Janashia Phil Kurt", the entire list
		//Index exchanged to underscore (_) to know it's there, we're just not using it as a var
		//strings.Fields = Splits "Nicole Smith" to become "Nicole","Smith"
		//We want the first element of the slice alone= "Nicole" Hence use of [0]
		//We get a list of firstNames (Nicole,etc) from all firstName.
		//append adds an element to a slice. firstNames is a slice. firstName is the added element to the slice. Bookings was also a slice, we added firstName&lastName
		//as elements to it.

		if remainingTickets == 0 {
			//end program
			fmt.Println("SOLD OUT!")
			break
		}

	}

}
