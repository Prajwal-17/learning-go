package main

import "fmt"

// &  address operator - to create a pointer
// * dereference operator - to access modify the value

func main() {
	score := 300
	fmt.Println("score:", score)

	p := &score
	fmt.Println("memory address of original score variable:", p)
	fmt.Println("value address of pointer:", *p)
	fmt.Println("memory address of pointer p:", &p)

	*p = 500
	fmt.Println("score value:", score)
	fmt.Println("new value of pointer p:", *p)
}
