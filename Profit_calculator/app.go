package main

import "fmt"

func main() {
	// var expenses float64
	// var taxRate float64
	// var name string

	// Input from user
	name := outputText("Enter Name: ")
	revenue := outputText("Enter Revenue: ")
	expenses := outputText("Enter Expenses: ")
	taxRate := outputText("Enter Tax rate: ")

	// Calculations
	ebt := revenue - expenses          
	profit := ebt * (1 - taxRate/100) 
	ratio := ebt / profit             

	// Output
	myName :=fmt.Sprintf("Name: %s", name)
	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio (EBT/Profit): %.2f\n", ratio)
	fmt.Println(myName)
}
func outputText(text string)float64  {
	var userOutput float64
	fmt.Print(text)
	fmt.Scan(&userOutput)
	return userOutput
}