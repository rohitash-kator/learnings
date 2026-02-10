package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank!")

	fmt.Println("Reach us 24/7:", randomdata.PhoneNumber())

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
		// panic("Can't continue, sorry!")
	}

	for {
		displayMenuOptions()

		var choice int

		fmt.Print("Your choice: ")
		fmt.Scan((&choice))

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("How much money do you want to deposit: ")
			var depositAmount float64
			fmt.Scan((&depositAmount))

			if depositAmount <= 0 {
				fmt.Println("Invalid amount entered. Amount must be greater than 0.")
				continue
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount

			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatFromFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("How much money do you want to withdraw: ")
			var withdrawalAmount float64
			fmt.Scan((&withdrawalAmount))

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount entered. Amount must be between 0.")
				// return
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount entered. You can't withdraw more than you have.")
				// return
				continue
			}

			accountBalance -= withdrawalAmount // accountBalance = accountBalance - withdrawalAmount

			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatFromFile(accountBalance, accountBalanceFile)

		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
			// break
		default:
			fmt.Println("Invalid choice...")
			continue
		}

	}

}
