package main

import "fmt"

func main() {

	var taxrate float64

	// fmt.Print("Enter revenue: ")
	// fmt.Scanln(&revenue)
	// printRevenue()

	// fmt.Print("Enter expenses: ")
	// fmt.Scanln(&expenses)
	// printExpens()

	revenue := getUserInput("Enter Revenue : ")
	expenses := getUserInput("Enter Expense : ")
	// var profit = revenue - expenses
	// var earningAfterTax = profit * (1 - taxrate/100)
	profit, earningAfterTax := caluclationOfProfit(revenue, expenses, taxrate)
	fmt.Println("Profit:", profit)
	fmt.Println("Earnings after tax:", earningAfterTax)

}

func caluclationOfProfit(revenue, expenses, taxrate float64) (float64, float64) {
	var profit = revenue - expenses
	var earningAfterTax = profit * (1 - taxrate/100)
	return profit, earningAfterTax
}

func getUserInput(textInput string) float64 {
	var userInput float64
	fmt.Print(textInput)
	fmt.Scanln(&userInput)
	return userInput

}

// func printRevenue() float64 {
// 	fmt.Print("Enter revenue: ")
// 	fmt.Scanln(&revenue)
// 	return revenue
// }

// func printExpens() float64 {
// 	fmt.Print("Enter expenses: ")
// 	fmt.Scanln(&expenses)
// 	return expenses
// }
