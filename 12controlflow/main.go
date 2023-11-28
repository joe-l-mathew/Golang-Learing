package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("IF and Else in golang")

	var loginCount int = 9

	if loginCount < 10 {
		fmt.Println("Regular user")
	} else if loginCount == 10 {
		fmt.Println("Exactly 10 User")
	} else {
		fmt.Println("Watch out")
	}

	// read a value from terminal and print weather its even or odd

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the number to check even or odd: ")

	data, _ := reader.ReadString('\n')

	intData, err := strconv.Atoi(strings.TrimSpace(data))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	number := intData

	if number == 0 {
		fmt.Println("The number is neither odd nor even")
	} else if number%2 != 0 {
		fmt.Println("The number is odd")
	} else {
		fmt.Println("The number is even")
	}

	//declaring with if check

	if a := number; a < 4 {
		fmt.Println("the value a is less than 4")
	}
}
