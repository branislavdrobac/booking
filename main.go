// by defaut we need to define package/program name
package main

// importing needed packages, "fmt" for text formating, "strings" for strings processing
import (
	"fmt"
	"strings"
)

func main() {
	// general variables
	// var conferenceName = "Go Conference"
	conferenceName := "Go Conference" // cleaner code variable then above
	const conferenceTickets int = 50  // const variable is variable that cannot be changed
	var remainingTickets uint = 50    // uint is unsigned integer and cannot be negative
	bookings := []string{}
	// printing welcome message bla bla
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
	// begin loop for user input
	for {
		// defining input variables
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name and details - user input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Printf("Memory address location: %p\n", &firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Printf("Memory address location: %p\n", &lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Printf("Memory address location: %p\n", &email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		fmt.Printf("Memory address location: %p\n", &userTickets)
		// tickets math
		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			// dynamic lists in slice
			bookings = append(bookings, firstName+" "+lastName)
			/*
				// debug and print verbose array/slice with type and length
				fmt.Printf("The whole slice: %v\n", bookings)
				fmt.Printf("The first value: %v\n", bookings[0])
				fmt.Printf("Slice type: %T\n", bookings)
				fmt.Printf("Slice length: %v\n", len(bookings))
				fmt.Printf("Memory address location: %p\n", &bookings)
			*/
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmatiom email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			// loop inside loop to extract first name of ticket reservation
			firstNames := []string{}
			// when you have empty or undeclared variable use _ to let go compiler know that this was intentional
			for _, booking := range bookings { // here we changed variable from "index" to "_" to avoid compiler errors
				var names = strings.Fields(booking)
				// var firstName = names[0] - pro tip - use index of slice/array right away like in below names[0], cleaner code
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break // end of program
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}
	// end of loop for user input
}

// end of program
