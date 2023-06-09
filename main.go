package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings []string

const conferenceTickets = 50

func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	/*
		We can declare variables using := syntax as shown below

		testVariable1 := "testVariable-1"
		fmt.Printf("I am %v \n \n", testVariable1)
	*/

	/*
			To display the datatype of a variable, use %T placeholder

		fmt.Printf("conferenceRickets is %T \n", conferenceTickets)
		fmt.Printf("conferenceName is %T \n", conferenceName)
		fmt.Printf("remainingTickets is %T \n \n", remainingTickets)

	*/

	/*
		Printf is used to display formatted text. For example adding a placeholder for a variable as shown below.
		%v should be used
		and the character for nextline is \n
	*/
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)

	/*
		Loops
	*/

	for {
		/*
		*	checking validity of inputs using functions
		 */

		firstName, lastName, email, userTickets := userInput()
		isValidName, isValidEmail, isValidUserTickets := checkValidity(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email)
			/*
			*	call function print first names
			 */
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of our bookings: %v \n \n", firstNames)

			// noTicketsRemain := remainingTickets == 0
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email you have entered is invalid!")
			}
			if !isValidUserTickets {
				fmt.Println("The tickets you have entered is invalid")
			}
		}

	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n \n")

}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func checkValidity(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	// In GO you can return more than one value to the calling function
	return isValidName, isValidEmail, isValidUserTickets
}

func userInput() (string, string, string, uint) {
	/*
		Arrays and Slices
		Slices are defined similar to arrays but only change is size definition is not given.
	*/
	// var bookings [50]string

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// User Inputs

	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last Name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email:")
	fmt.Scan(&email)
	fmt.Println("Enter no of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("The Array type: %T\n", bookings)
	// fmt.Printf("Array Length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive the tickets in your email address %v\n \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining\n", remainingTickets)
}
