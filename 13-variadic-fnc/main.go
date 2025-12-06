package main

import "fmt"

// Variadic func => a function that can take dynamic number of arguments of same type

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Variadic functions")
	result := sum(5, 34, 23)
	fmt.Println("Sum", result)
	nums := []int{4, 5, 6}
	fmt.Println("nums", sum(nums...))
}
