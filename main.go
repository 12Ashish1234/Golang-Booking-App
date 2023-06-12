package main

import (
	"fmt"
	"time"
)

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

const conferenceTickets = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

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

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := CheckValidity(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			/*
			*	call function print first names
			 */
			firstNames := getFirstNames()
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

			fmt.Println("Please try again!!! \n")
		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n \n")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// below method is for getting firstNames from normal strings
		// var names = strings.Fields(booking)

		// below method is for getting firstNames from maps
		// firstNames = append(firstNames, booking["firstName"])

		//below method is for getting firstNames from structs

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// func checkValidity(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

// 	// In GO you can return more than one value to the calling function
// 	return isValidName, isValidEmail, isValidUserTickets
// }

func getUserInput() (string, string, string, uint) {
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create map for a user
	// var userData = make(map[string]string)

	// create struct for userdata
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v\n", bookings)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("The Array type: %T\n", bookings)
	// fmt.Printf("Array Length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive the tickets in your email address %v\n \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining\n", remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###############")

}
