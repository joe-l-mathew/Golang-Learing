package main

import "fmt"

func main() {

	var mySlice []string = []string{
		"Joel", "Aju", "manuval", "Aswin", "Pooja",
	}

	//going to remove Manuval which is at index 2
	var indexToRemove int = 2
	mySlice = append(mySlice[:indexToRemove], mySlice[indexToRemove+1:]...)

	fmt.Println(mySlice)

	// scores := make([]int, 10)
	// scores = append(scores, scores...)
	// fmt.Println(scores)

}
