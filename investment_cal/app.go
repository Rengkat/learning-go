package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years float64 = 1000,10//must be of same type
	expectedReturnRate := 2.2 //go infer the type
	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter year: ")
	fmt.Scan(&years)
	fmt.Print("Expected Return: ")
	fmt.Scan(&expectedReturnRate)
	futureValue := investmentAmount * math.Pow( 1 + expectedReturnRate / 100,years)
	fmt.Println(futureValue)
}