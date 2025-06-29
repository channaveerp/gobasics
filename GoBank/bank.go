package main

import (
	"fmt"
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

	if choice == 1 {
		fmt.Println("Your available Balance is ", avaibelance)
	} else if choice == 2 {
		fmt.Scan(&addmoney)
		if addmoney <= 0 {
			fmt.Println("Invallid Amount! amount must be greater then 0")
			return
		}
		fmt.Println("your money added successfulley ", addmoney)
		avaibelance += addmoney
		fmt.Println("Now Total Available Bance is ", avaibelance)
	} else if choice == 3 {
		fmt.Scan(&widraw)
		if widraw <= 0 {
			fmt.Println("Invallid Amount! amount must be greater then 0")
			return
		}

		if widraw > avaibelance {
			fmt.Println("Abe sale Garib Utna paisa nahi hi tere pas jitna tu mangrah Hi")
			return
		}
		fmt.Println("Your Widra money is ", widraw)
		fmt.Println("Now Your Available Bance is ", avaibelance-widraw)
	} else {
		fmt.Println("GoodBye!")

	}

}
