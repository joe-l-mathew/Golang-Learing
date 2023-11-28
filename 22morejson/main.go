package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"name"`
	Platform string `json:"website"`
	Price    int
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcom to JSON video")
	EncodeJson()
}

func EncodeJson() {
	myCourses := []courses{
		{Name: "React JS", Platform: "Youtube", Price: 00, Password: "mySecretPassword1", Tags: []string{"JS", "Web"}},
		{Name: "Angular JS", Platform: "Youtube", Price: 00, Password: "mySecretPassword2"},
		{Name: "Spring Boot", Platform: "Udemy", Price: 499, Password: "mySecretPassword3", Tags: []string{"JAVA", "Web", "Beens"}},
	}
	jsonData, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(jsonData)
	}
	fmt.Println(string(jsonData))

}
