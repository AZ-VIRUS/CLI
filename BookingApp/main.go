package main

import (
	"BookingApp/helper"
	"fmt"
	"time"
)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
	selectedAnime  string
}

var bookingName = "Anime Booking"

const totalTickets = 100

var availableTickets uint = 50
var bookings = make([]UserData, 0)

// Anime list
var animes = []string{
	"Naruto",
	"Attack on Titan",
	"One Piece",
	"Demon Slayer",
	"Jujutsu Kaisen",
}

func main() {

	greetUser()

	for {

		showAnimeList()

		selectedAnime := chooseAnime()

		firstName, lastName, email, userTicket := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber :=
			helper.ValidateUserInput(firstName, lastName, email, userTicket, availableTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, email, userTicket, selectedAnime)
			go sendTicket(userTicket, firstName, lastName, email)

			getFirstNames()

			showBookings()

			if availableTickets == 0 {
				fmt.Println("Tickets are sold out!")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("Name must be at least 2 characters.")
			}

			if !isValidEmail {
				fmt.Println("Email must contain @ symbol.")
			}

			if !isValidTicketNumber {
				fmt.Println("Invalid ticket amount.")
			}

			fmt.Println("Please try again.")
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v!\n", bookingName)
	fmt.Printf("Remaining tickets: %v / %v\n", availableTickets, totalTickets)
}

func showAnimeList() {

	fmt.Println("\nAvailable Anime Shows:")

	for i, anime := range animes {
		fmt.Printf("%v. %v\n", i+1, anime)
	}
}

func chooseAnime() string {

	var animeChoice int

	fmt.Print("Select anime number: ")
	fmt.Scanln(&animeChoice)

	if animeChoice < 1 || animeChoice > len(animes) {
		fmt.Println("Invalid selection")
		return chooseAnime()
	}

	return animes[animeChoice-1]
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Print("First Name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Email: ")
	fmt.Scanln(&email)

	fmt.Print("Number of tickets: ")
	fmt.Scanln(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(firstName string, lastName string, email string, userTicket uint, selectedAnime string) {

	availableTickets -= userTicket

	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: userTicket,
		selectedAnime:  selectedAnime,
	}

	bookings = append(bookings, userData)

	fmt.Printf("\n%v %v booked %v ticket(s) for %v\n",
		firstName, lastName, userTicket, selectedAnime)

	fmt.Printf("Remaining tickets: %v\n", availableTickets)
}

func getFirstNames() {

	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	fmt.Println("Booked users:", firstNames)
}

func showBookings() {

	if len(bookings) == 0 {
		fmt.Println("No bookings yet.")
		return
	}

	fmt.Println("\nAll Bookings:")

	for i, booking := range bookings {

		fmt.Printf("%v. %v %v | %v | %v ticket(s) | Anime: %v\n",
			i+1,
			booking.firstName,
			booking.lastName,
			booking.email,
			booking.numberOfTicket,
			booking.selectedAnime,
		)
	}
}
func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("***********")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("Ticket sent successfully!")

}
