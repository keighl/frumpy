# Frumpy

A golang library to filter unwanted keys from an arbitrary JSON objects.

```go

jsonIn := []byte(`{...}`)

jsonOut, err := frumpy.FilterJSON(jsonIn, "id", "author.id", "comments.id", "comments.user.id")
if err != nil {
	panic(err)
}

// Look, mom! No more ids.
fmt.Println(string(jsonOut))

```
