package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().Unix())

	diceValue := rand.Intn(6) + 1

	switch diceValue {
	case 1:
		fmt.Println("The number is One")
	case 2:
		fmt.Println("The number is Two")
		fallthrough
	case 3:
		fmt.Println("The number is three")
	case 4:
		fmt.Println("The number is Four")
	case 5:
		fmt.Println("The number is Five")
	case 6:
		fmt.Println("The number is Siz, and you can throw once more")
	default:
		fmt.Println("What was that")
	}
}
