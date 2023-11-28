package main

import "fmt"

func main() {
	var myarr [4]string = [4]string{
		"Element1",
		"Element 2",
		"Element 3",
	}
	// myarr[0] = "Element 1"
	// myarr[1] = "Element 2"
	// myarr[2] = "Element 3"
	// myarr[3] = "Element 4"
	fmt.Printf("The type of myarr is: %T\n", myarr)
	fmt.Println(myarr)
}
