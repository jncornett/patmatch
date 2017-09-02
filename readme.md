[![GoDoc](https://godoc.org/github.com/jncornett/patmatch?status.svg)](https://godoc.org/github.com/jncornett/patmatch)

The reverse `printf`. This library is for doing things like:

```go
t, _ := patmatch.Parse("Hello, %(name)s")
unprintf := t.Apply("Hello, world")
fmt.Println("The greeting was meant for", unprintf["name"])
```
