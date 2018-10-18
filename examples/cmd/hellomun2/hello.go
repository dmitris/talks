package main

import (
	"fmt"
	"os"

	"github.com/dmitris/talks/examples/hellomun"
	"github.com/dmitris/talks/examples/hellonew"
)

func main() {
	audience := "Munich Gophers"
	if len(os.Args) > 1 {
		audience = os.Args[1]
	}

	greeting := hellomun.Greeting(audience)
	fmt.Println(greeting)
	localGreeting := hellonew.LocalGreeting()
	fmt.Println(localGreeting)
}
