package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const js = `{
	"name": "Dmitrij",
	"age": 27
}`

const complex_js = `
	[
		{
			"name": "Dmitrij",
			"age": 27
		},
		{
			"name": "Alex",
			"age": 30
		}
	]`

func main() {
	var user User

	// from json to struct
	err := json.Unmarshal([]byte(js), &user)

	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	// from struct to json
	user_json, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(user_json))

	fmt.Println()
	// complex_js

	var users []User

	err = json.Unmarshal([]byte(complex_js), &users)

	if err != nil {
		panic(err)
	}

	fmt.Println(users)

	users_js, err := json.Marshal(users)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(users_js))
}
