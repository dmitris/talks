package main

import (
	"fmt"
	"os"

	"github.com/dmitris/talks/examples/hellomun"
)

func main() {
	audience := "Munich Gophers"
	if len(os.Args) > 1 {
		audience = os.Args[1]
	}

	greeting := hellomun.Greeting(audience)
	fmt.Println(greeting)
}
