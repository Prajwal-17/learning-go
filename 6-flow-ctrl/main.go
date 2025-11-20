package main

import "fmt"

func main() {
	foreg1()
	while()
}

func foreg1() {
	sum := 0
	for i := 0; i < 20; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func while() {
	sum := 1
	for sum < 10 {
		sum += sum
		// sum++
	}
	fmt.Println(sum)
	// infinite loop
	// for {

	// }
}
