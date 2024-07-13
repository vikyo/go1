package main

import (
	"errors"
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

	var appUser, err = newUser(firstNameInput, lastNameInput, birthdateInput)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func newUser(firstNameInput, lastNameInput, birthdateInput string) (*user, error) {
	if firstNameInput == "" || lastNameInput == "" || birthdateInput == "" {
		return nil, errors.New("invalid inputs")
	}
	return &user{
		firstName: firstNameInput,
		lastName:  lastNameInput,
		birthdate: birthdateInput,
		createdAt: time.Now(),
	}, nil
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
	fmt.Scanln(&value)
	return value
}
