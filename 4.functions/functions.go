package main

import "fmt"

// if both params are same type then -> x, y int
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("result is", add(2, 3))
}
