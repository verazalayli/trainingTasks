package task2_jsonParser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func Task2() {
	users := readJson("task2_jsonParser/file.json")
	youngestUser := User{}
	for _, user := range users {
		if youngestUser.Name == "" {
			youngestUser = user
		}
		if user.Age <= youngestUser.Age {
			youngestUser = user
		}
	}
	fmt.Printf("Youngest: %s (%d)\n", youngestUser.Name, youngestUser.Age)
}

func readJson(str string) []User {
	data, err := os.ReadFile(str)
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatal(err)
	}
	return users
}
