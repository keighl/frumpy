package main

import (
	"fmt"

	"github.com/keighl/frumpy"
)

func main() {
	// Let's take all the IDs out the `jsonIn` payload

	var badKeys = []string{
		// Top level ID
		"id",
		// Nested object
		"brand.id",
		// Double nested object!
		"brand.ceo.id",
		// Nested array of comments
		"comments.id",
		// Nested object inside nested array of comments!
		"comments.user.id",
	}

	jsonOut, err := frumpy.FilterJSON(jsonIn, badKeys...)
	if err != nil {
		panic(err)
	}

	// Look, mom! No more ids.
	fmt.Println(string(jsonOut))
}

var jsonIn = []byte(`{
	"id": "XXXXXX",
	"name": "Cheese Sticks",
	"price": 59.00,
	"description": "",
	"brand": {
		"id": "nabisco",
		"name": "Nabisco",
		"ceo": {
			"name": "Carl",
			"id": "GUSTO"
		}
	},
	"comments": [
		{
			"id": "XXXXX",
			"body": "lorem",
			"user": {
				"name": "Bob",
				"id": "55555"
			}
		},
		{
			"id": "YYYYY",
			"body": "lorem",
			"user": {
				"name": "Jane",
				"id": "66666"
			}
		}
	]
}`)
