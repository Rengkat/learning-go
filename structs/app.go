package main

import (
	"fmt"
	"github/Rengkat/structs/user"
)

func main() {
	var name string
	var age int
	var height float64

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter age: ")
	fmt.Scanln(&age)
	fmt.Print("Enter height: ")
	fmt.Scanln(&height)

	// Create a User instance
	user, err := user.NewUser(name, age, height)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Call methods
	fmt.Println(user.Desc()) // Print description
	user.ClearName()         // Clear the name
	fmt.Println(user.Desc()) // Print updated description
}