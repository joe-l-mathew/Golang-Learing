package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	// weekdays := []string{
	// 	"Sunday",
	// 	"Monday",
	// 	"Tuesday",
	// 	"Wedensday",
	// 	"Thursday",
	// 	"Friday",
	// 	"Saturday",
	// }

	// for i := range weekdays {
	// 	fmt.Println(weekdays[i])
	// }

	// for i, v := range weekdays {
	// 	fmt.Printf("The weeday %v is at the index %v\n", v, i)
	// }

	// for i := 0; i < len(weekdays); i++ {
	// 	fmt.Println(weekdays[i])
	// }

	limitVal := 10
	initValue := 1

	for initValue <= limitVal {

		if initValue == 2 {
			goto sudo // will go to the sudo label
		}
		if initValue == 5 {
			// break
			initValue++
			continue
		}
		fmt.Println("The value is :", initValue)
		initValue++

	}

	//goto in golang
sudo:
	fmt.Println("Reached the label sudo")

}
