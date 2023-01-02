package main

import (
	"fmt"
	"golang/greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println(greetings.Hello("me"))
	fmt.Println(quote.Go())
}
