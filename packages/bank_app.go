package main

import (
	"fmt"

	"example.com/bank-app/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFile = "balance.txt"

func main() {
	currentBalance, err := fileops.ReadFloatValueFromFile(balanceFile)

	if err != nil {
		fmt.Println(err)
		fmt.Print("----------")
		panic("Can't continue")
	}

	var amount float64
	var choice int
	fmt.Println("Welcome to bank!")
	fmt.Println("Reach us at: ", randomdata.PhoneNumber())

	for {
		getOptions()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			getBalance(currentBalance)
		case 2:
			fmt.Println("Enter Amount: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Inavalid amount")
				continue
			}
			currentBalance += amount
			getBalance(currentBalance)
			fileops.WriteFloatValueToFile(currentBalance, balanceFile)
		case 3:
			fmt.Println("Enter Amount: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Inavalid amount")
				continue
			}
			if amount > currentBalance {
				fmt.Println("Inavalid amount")
				continue
			}
			currentBalance -= amount
			getBalance(currentBalance)
			fileops.WriteFloatValueToFile(currentBalance, balanceFile)
		case 4:
			fmt.Println("Goodbye")
			return
		default:
			fmt.Println("Wrong choice")
		}
	}

}

func getBalance(balance float64) {
	fmt.Println(balance)
}
