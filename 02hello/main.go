package main

import "fmt"


func main() {
	var username string = "Joel Mathew"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var smallInt uint
	fmt.Println(smallInt)

	var name string
	fmt.Printf("The name is havn=ing the value:%v", name)

	var smallFloat float32
	fmt.Printf("The variable value is: %v, type is: %T\n", smallFloat, smallFloat)

	var url = "bookmywash.com"
	fmt.Printf("The site url is: %v\n", url)

	url2 := "url2"
	fmt.Println(url2)
}
