package main

import "fmt"

func main() {

	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

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
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here \n \n")

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

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive the tickets in your email address %v\n \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining\n", remainingTickets)

}
