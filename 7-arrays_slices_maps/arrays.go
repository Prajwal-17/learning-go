package main

import "fmt"

// array is fixed in size
func arrays() {
	var a [5]int                  //initialize an array
	fmt.Println("empty array", a) // [0 0 0 0 0]
	a[4] = 100
	fmt.Println("set:", a)

	var b [5]int = [5]int{1, 2, 3, 4, 5}
	for i := range b {
		fmt.Println("loop array", i)
	}
	fmt.Println("array", b)

	// 2-D array
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println(matrix) // [[1 2 3] [1 2 3] [1 2 3]]

}
