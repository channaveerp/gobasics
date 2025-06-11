package main

import (
	"fmt"
	"os"
)

func main() {
	var avaibelance float64 = 1000
	var addmoney float64
	var widraw float64

	fmt.Println("Wel come Go Bank!")
	fmt.Println("What You want to do?")
	fmt.Println("Select Your Choices")
	fmt.Println("1. Check Balace")
	fmt.Println("2. Deposite Money")
	fmt.Println("3. Widraw Money")
	fmt.Println("4. Exit")

	// take user choice
	var choice int
	fmt.Scan(&choice)
	fmt.Println("your choice is", choice)

	// check avaibale balance
	var wantsTocheckBalnce = choice == 1
	if wantsTocheckBalnce {
		fmt.Println("Your available Balance is ", avaibelance)
	}
	// add money
	var wantsToAddmoney = choice == 2
	if wantsToAddmoney {
		fmt.Scan(&addmoney)
		fmt.Println("your money added successfulley ", addmoney)
		fmt.Println("Now Total Available Bance is ", avaibelance+addmoney)
	}

	//  widra money
	var wantsTowidrawMoney = choice == 3
	if wantsTowidrawMoney {
		fmt.Scan(&widraw)
		fmt.Println("Your Widra money is ", widraw)
		fmt.Println("Now Your Available Bance is ", avaibelance-widraw)
	}
	// exit
	var wantsToExit = choice == 4
	if wantsToExit {
		fmt.Println("Your selected choice is", choice)
		os.Exit(0)

	}

}
