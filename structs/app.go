package main

import (
	"fmt"
	"structs/user"
)

func main() {
	firstNameInput := getUserData("Please enter your first name: ")
	lastNameInput := getUserData("Please enter your last name: ")
	birthdateInput := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.UserModel
	var err error
	var adminUser *user.AdminUserModel

	appUser, err = user.NewUser(firstNameInput, lastNameInput, birthdateInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	adminUser, err = user.NewAdminUser("test@test.com", "1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	adminUser.OutputUserDetails()
	adminUser.ClearUserName()
	adminUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
