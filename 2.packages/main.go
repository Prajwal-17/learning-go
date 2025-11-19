package main

import (
	"fmt"

	"math/rand"

	"rsc.io/quote" // external package
)

func main() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
	fmt.Println(rand.Int31())
}
