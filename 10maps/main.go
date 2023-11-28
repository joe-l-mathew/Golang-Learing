package main

import "fmt"

func main() {
	var languages = make(map[string]string)
	languages["JS"] = "javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	// delete(languages, "JS")
	// fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("The key is %v and the value is %v\n", key, value)
	}
}
