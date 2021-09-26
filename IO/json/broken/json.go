package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const js = `
	[{
		"name":"Axel",
		"lastname":"Fooley"
	},
	{
		"name":"Tim",
		"lastname":"Burton"
	},
	{
		"name":"Tim",
		"lastname":"Burton"
`

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

func main() {
	users := make([]User, 0)

	r := strings.NewReader(js)
	dec := json.NewDecoder(r)

	for {
		token, err := dec.Token()

		if err != nil {
			break
		}

		if token == nil {
			break
		}

		switch tp := token.(type) {
		case json.Delim:
			str := tp.String()
			if str == "[" || str == "{" {
				for dec.More() {
					u := User{}
					err := dec.Decode(&u)

					if err == nil {
						users = append(users, u)
					} else {
						break
					}
				}
			}
		}
	}

	fmt.Println(users)
}
