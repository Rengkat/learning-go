package main

import "fmt"

type User struct {
	name   string
	age    int
	height float64
}
// desc is a method on the User type
func (u User) desc() {
	fmt.Printf("My name is %s. I am %d years old and %.2f meters tall.\n", u.name, u.age, u.height)
}
func (u *User) clearName()  {
	u.name=""
}
func main() {
	var name string
	var age int
	var height float64

	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	fmt.Print("Enter age: ")
	fmt.Scan(&age)
	fmt.Print("Enter height: ")
	fmt.Scan(&height)

	// Create a User instance with the input values
	user := User{
		name:   name,
		age:    age,
		height: height,
	}

	// Call the desc method on the User instance
	user.desc()
	user.clearName()
	user.desc()

}

