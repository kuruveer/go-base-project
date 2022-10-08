package main

import (
	"fmt"
	"base/helper"
	"sync"
	"time"
)

const flightName = "Flight A654"
const flightTicketCount = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main(){
	greetingMessage()

	for{
		firstName, lastName, email, numberOfTickets := getUserInputs()

		isValidName, isValidEmail, isValidNumber := helper.ValidateUserInput(firstName, lastName, email, numberOfTickets, remainingTickets)

		if isValidName && isValidEmail && isValidNumber {
			bookTicket(firstName, lastName, email, numberOfTickets)

			firstNames := getFirstName()

			fmt.Printf("Passengers List %v\n", firstNames)

      wg.Add(1)
      go sendTicket(userTickets, firstName, lastName, email)
  
			if remainingTickets == 0 {
				break
			}
		}else{
			if !isValidName {
				fmt.Println("first or last name you enter is short")
			}
			if !isValidEmail {
				fmt.Println("entered email is invalid")
			}
			if !isValidNumber {
				fmt.Println("Inavlid number for tickets")
			}
		}
	}
  wg.Wait()
}

func greetingMessage(){
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("Welcome to Veer Flights")
	fmt.Printf("We have a %v tickets and %v are still avaliable for booking \n", flightTicketCount, remainingTickets)
	fmt.Println("Get your tickets know")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@")
}

func getUserInputs() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var numberOfTickets uint

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&numberOfTickets)

	return firstName, lastName, email, numberOfTickets
}

func bookTicket(firstName string, lastName string, email string , numberOfTickets uint){
	remainingTickets = remainingTickets - numberOfTickets

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: numberOfTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v, for booking %v tickets in our flight %v \n", firstName, lastName, numberOfTickets, flightName)
	fmt.Printf("%v tickets remaining of %v \n", remainingTickets, flightTicketCount)
}

func getFirstName() []string {
	firstNames := []string{}

	for _, booking := range(bookings){
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
