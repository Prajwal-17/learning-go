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

// --- OUTPUT ---
// score: 300
// memory address of original score variable: 0xc000012130
// value address of pointer: 300
// memory address of pointer p: 0xc00005e048
// score value: 500
// new value of pointer p: 500
