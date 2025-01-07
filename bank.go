package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFileName = "balance.txt"

func readBalanceFromFile() (float64, error) {
	content, err := os.ReadFile(accountBalanceFileName)

	if err != nil {
		return 0, errors.New("Balance file not found")
	}

	balancetext := string(content)
	balance, error := strconv.ParseFloat(balancetext, 64)

	if error != nil {
		return 0, errors.New("Failed to parse balance")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balancetext := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFileName, []byte(balancetext), 0644)
}

func main() {
	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR: ", err)
		fmt.Println("-----------------")
	}

	fmt.Println("Welcome to Go bank")

	for {
		presentOptions()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Println("How much would you like to deposit?")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! new amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Println("How much would you like to withdraw?")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Withdrawal amount must be greater than 0")
				return
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Withdrawal amount is greater than your balance")
				return
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! new amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing Go bank")
			return
		}

		fmt.Println("Your choice is", choice)
	}
}
