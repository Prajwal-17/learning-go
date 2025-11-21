package main

import "fmt"

// A slice is a flexible, dynamically-sized abstraction over an array. Slices provide a way to work with subsets of arrays.
// a slice are reference types, meaning they reference an underlying array
// A slice has three components: a pointer to the array, a length, and a capacity.
func slices() {
	var users []string
	println(users) // [0/0]0x0

	var numbers []int = []int{1, 2, 3, 4}
	fmt.Println(numbers)

	// make is used to initialize new array|slice|maps => make(type, length, capacity)
	// length - no of elements in a slice
	// capacity - the max no of elements that a slice can hold without requiring memory allocation

	// initialize without default value
	var names []string = make([]string, 3)
	fmt.Println("names", cap(names)) // capacity
	fmt.Println("names", len(names)) // length
	fmt.Println("names", names)

	// copy by reference
	var names_1 []string = []string{"prajwal", "kirat", "rahul"}
	var names_2 = names_1 // copy by ref
	names_2[2] = "rohit"
	fmt.Println("copy by reference", names_2)

	// copy by value
	var clg_names_1 []string = []string{"bit", "vit", "rv"}
	var clg_names_2 = make([]string, len(clg_names_1))
	copy(clg_names_2, clg_names_1)
	clg_names_2[2] = "reva"
	fmt.Println("copy by value", clg_names_2)
}
