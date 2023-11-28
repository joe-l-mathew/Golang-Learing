package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	var email = "joelmathew357"
	joel := User{
		Name:   "Joel",
		Email:  &email,
		Status: true,
		Age:    22,
	}
	fmt.Println(joel)
	fmt.Println("Email id is:" + *joel.Email)
	fmt.Printf("The values are: %+v\n", joel)
}

type User struct {
	Name   string
	Email  *string
	Status bool
	Age    int
}
