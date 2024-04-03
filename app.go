package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("--------------------------")
		// panic("App go-basic-bank is crashing!")
	}

	fmt.Println("Welcome to GoBank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit balance")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Deposit amount must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! Your balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Print("Your withdrawal: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Withdrawal amount must be greater than 0.", accountBalance)
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. Withdrawal amount should not be greater than account balance.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! Your balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using our bank!")
			return
		}
	}

}
