package patmatch

import "fmt"

func Example() {
	templ, err := ParseFlags("twas %(time)s, and the %s toves did %s and %s in the %(place)s", 0)
	if err == nil {
		panic(err)
	}

	fmt.Println(templ.Expr())

	matches := templ.Apply("twas brillig, and the slithy toves did gyre and gimble in the wabe")
	fmt.Println(matches)
}

func Example_helloWorld() {
	templ, err := ParseFlags("Hello, %(name)s", 0)
	if err != nil {
		panic(err)
	}

	matches := templ.Apply("Hello, world")
	fmt.Println("the subject of the greeting was", matches["name"])
}
