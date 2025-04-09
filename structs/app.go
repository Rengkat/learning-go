package main

import (
	"errors"
	"fmt"
)

type User struct {
	name   string
	age    int
	height float64
}

// Using pointer receiver for consistency
func (u *User) desc() {
	fmt.Printf("My name is %s. I am %d years old and %.2f meters tall.\n", u.name, u.age, u.height)
}

func (u *User) clearName() {
	u.name = ""
}

func newUser(name string, age int, height float64) (*User, error) {
	if name == "" || age <= 0 || height <= 0 {
		return nil, errors.New("please enter all fields with valid values")
	}
	return &User{
		name:   name,
		age:    age,
		height: height,
	}, nil
}

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

	// Create a User instance using the newUser function
	user, err := newUser(name, age, height)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Call methods on the User instance
	user.desc()
	user.clearName()
	user.desc()
}