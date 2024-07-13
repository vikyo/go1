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

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func (u user) outputUserDetails() {
	fmt.Println("user in outputUserDetails:", u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *user) clearUserName() {
	fmt.Println("clearing name")
	u.firstName = ""
	u.lastName = ""
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
