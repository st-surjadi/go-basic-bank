package main

import (
	"fmt"

	"example.com/go-basic-bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "_balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("--------------------------")
		// panic("App go-basic-bank is crashing!")
	}

	fmt.Println("Welcome to GoBank!")
	fmt.Println("Reach us 24/7 on ", randomdata.PhoneNumber())

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using our bank!")
			return
		}
	}

}
