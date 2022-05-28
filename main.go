package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJSON := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name: "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJSON), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)
}

// import (
// 	"LearnPackages/helpers"
// 	"log"
// )

// const numPool = 1000

// func CalculateValue(intChan chan int) {
// 	randomNumber := helpers.RandomNumber(numPool)
// 	intChan <- randomNumber
// }

// func main() {
// 	intChan := make(chan int)
// 	defer close(intChan)

// 	go CalculateValue(intChan)

// 	num := <-intChan
// 	log.Println(num)
// }
