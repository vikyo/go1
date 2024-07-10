package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func main() {
	currentBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println(err)
		fmt.Print("----------")
		panic("Can't continue")
	}

	var amount float64
	var choice int
	fmt.Println("Welcome to COM bank!")

	for {
		fmt.Println("What you want to do ?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(currentBalance)
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
			writeBalanceToFile(currentBalance)
		default:
			fmt.Println("Goodbye")
			return
		}
	}

}

func getBalance(balance float64) {
	fmt.Println(balance)
}

func writeBalanceToFile(balance float64) {
	os.WriteFile(balanceFile, []byte(fmt.Sprint(balance)), 0644)
}

func readBalanceFromFile() (float64, error) {
	arg, err := os.ReadFile(balanceFile)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}
	balanceText := string(arg)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse amount from file")
	}

	return balance, nil
}
