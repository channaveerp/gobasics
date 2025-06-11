package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmout float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmout)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the year: ")
	fmt.Scan(&years)

	// finalAmount := investmentAmout * math.Pow(1+expectedReturnRate/100, years)
	// futureValue := finalAmount / math.Pow(1+inflationRate/100, years)
	finalAmount, futureValue := calculateFutureValue(investmentAmout, expectedReturnRate, years)

	fmt.Println(finalAmount)
	fmt.Println(futureValue)

}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	finalAmount := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futireValue := finalAmount / math.Pow(1+6.5/100, years)
	return finalAmount, futireValue

}
