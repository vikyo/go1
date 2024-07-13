package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	firstNameInput := getUserData("Please enter your first name: ")
	lastNameInput := getUserData("Please enter your last name: ")
	birthdateInput := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser = user{
		firstName: firstNameInput,
		lastName:  lastNameInput,
		birthdate: birthdateInput,
		createdAt: time.Now(),
	}
	fmt.Println("appUser in main: ", &appUser)
	fmt.Printf("Address of appUser in main: %p\n", &appUser)
	outputUserDetails(&appUser)
	fmt.Println("appUser in main:", appUser.firstName, appUser.lastName, appUser.birthdate, appUser.createdAt)
}

func outputUserDetails(u *user) {
	fmt.Println("App user in outputUserDetails: ", *u)
	fmt.Printf("Address of user in outputUserDetails: %p\n", u)
	u.firstName = "test"
	fmt.Println("user in outputUserDetails:", u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
