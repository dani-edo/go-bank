package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go bank")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Println("How much would you like to deposit?")
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Deposit amount must be greater than 0")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Balance updated! new amount: ", accountBalance)
	} else if choice == 3 {
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
	} else {
		fmt.Println("Goodbye!")
	}

	fmt.Println("Your choice is", choice)
}
