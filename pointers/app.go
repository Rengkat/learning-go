package main

import "fmt"
func add(a *int)  {
	*a++
}

func main() {
	num := 5
	mAddress :=&num
	fmt.Println(mAddress)
	fmt.Println(num)
	*mAddress = 10
	fmt.Println(*mAddress)
	add(&num)
	fmt.Println(num)
	//structs
		type Person struct{
		Name string
		Age int
	}
	person1 := Person{"Alexander Rengkat",30}
	fmt.Println(person1.Name)
	person2 := Person{
		Name: "John Doe",
		Age: 50,
	}
	fmt.Println(person2.Age)
}
//struct
