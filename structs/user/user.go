package user

import (
	"errors"
	"fmt"
)

type User struct {
	Name   string
	Age    int
	Height float64
}

// Exported method (note the uppercase 'D' in Desc)
func (u *User) Desc() string {
	return fmt.Sprintf(
		"My name is %s. I am %d years old and %.2f meters tall.",
		u.Name, u.Age, u.Height,
	)
}

// Exported method to clear the name
func (u *User) ClearName() {
	u.Name = ""
}

// Exported constructor function
func NewUser(name string, age int, height float64) (*User, error) {
	if name == "" || age <= 0 || height <= 0 {
		return nil, errors.New("please enter all fields with valid values")
	}
	return &User{
		Name:   name,
		Age:    age,
		Height: height,
	}, nil
}