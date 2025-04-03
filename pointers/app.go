package main

import "fmt"

// Person struct
type Person struct {
	Name string
	Age  int
}

// Employee struct
type Employee struct {
	ID        int
	Name      string
	Position  string
	Salary    float64
	IsManager bool
	Projects  []string
}

// Company struct
type Company struct {
	Name     string
	Employee Employee // Changed to uppercase to make it exported
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}
	type (
    Age      int
    Email    string
    PersonID string
)

type Person2 struct {
    ID    PersonID
    Name  string
    Age   Age
    Email Email
}
func (me Person2) Greeting() string {
	return "Good morning " + me.Name
}

func main() {
	// Initializing Person structs
	person1 := Person{"Alexander Rengkat", 30}
	fmt.Println(person1.Name) // "Alexander Rengkat"

	person2 := Person{
		Name: "John Doe",
		Age:  50,
	}
	fmt.Println(person2.Age) // 50

	var person3 Person
	person3.Name = "Alex 0"
	person3.Age = 24

	// Initializing Company with Employee
	company1 := Company{
		Name: "John Solution",
		Employee: Employee{
			ID:        23,
			Name:      "jina",
			Position:  "senior dev",
			Salary:    1234.6,
			IsManager: false,
			Projects:  []string{"song"},
		},
	}

	fmt.Println(company1.Name)              // "John Solution"
	fmt.Println(company1.Employee.Name)     // "jina"
	fmt.Println(company1.Employee.Projects) // ["song"]

	// Using the Greet method
	person := Person{Name: "Frank"}
	fmt.Println(person.Greet()) // 


var friend = Person2{"234","LEXANDER", 23, "alexan@gmail.com"}
fmt.Println(friend.Greeting())
}