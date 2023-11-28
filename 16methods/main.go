package main

import "fmt"

//functions in classes are regulerly known as methods,
//In golang functions in structs are known as methods

func main() {

	var joel User = User{"Joel", 22, "joelmathew357@gmail.com", true}
	fmt.Println(joel)
	joel.ManipulateEmail()
	fmt.Printf("The real values including keys %+v\n", joel)
	fmt.Println(joel.Email)
	fmt.Println(joel.GetUserStatus())
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) GetUserStatus() bool {
	return u.Status
}

func (u User) ManipulateEmail() {
	u.Email = "I had manipulated"
	fmt.Println(u.Email)
}
