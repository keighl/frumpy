package frumpy

import (
	"testing"
	"strings"
)

var j = []byte(`{
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

var badKeys = []string{
	"id",
	"brand.id",
	"brand.ceo.id",
	"comments.id",
	"comments.user.id",
}

func TestFilterJSON(t *testing.T) {
	b, err := FilterJSON(j, badKeys...)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if strings.Contains(string(b), `"id"`) {
		t.Fatalf(`Expected to strip all "id" keys. \n\n %s`, string(b))
	}
}

func BenchmarkFilterJSON(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		_, err := FilterJSON(j, badKeys...)
		if err != nil {
			b.Fatalf("error:", err)
		}
	}
}
