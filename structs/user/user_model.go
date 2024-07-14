package user

import (
	"errors"
	"fmt"
	"time"
)

type UserModel struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type AdminUserModel struct {
	email    string
	password string
	UserModel
}

func NewUser(firstNameInput, lastNameInput, birthdateInput string) (*UserModel, error) {
	if firstNameInput == "" || lastNameInput == "" || birthdateInput == "" {
		return nil, errors.New("invalid inputs")
	}
	return &UserModel{
		firstName: firstNameInput,
		lastName:  lastNameInput,
		birthdate: birthdateInput,
		createdAt: time.Now(),
	}, nil
}

func NewAdminUser(email, password string) (*AdminUserModel, error) {
	if email == "" {
		return nil, errors.New("invalid inputs")
	}
	return &AdminUserModel{
		email:    email,
		password: password,
		UserModel: UserModel{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "**",
			createdAt: time.Now(),
		},
	}, nil
}

func (u UserModel) OutputUserDetails() {
	fmt.Println("user in outputUserDetails:", u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *UserModel) ClearUserName() {
	fmt.Println("clearing name")
	u.firstName = ""
	u.lastName = ""
}
