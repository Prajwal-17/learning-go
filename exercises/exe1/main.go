package main

import "fmt"

func main() {
	// exe1
	// var a [5]int = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(sumofArray(a))

	// exe2
	score := 86
	grade := getgrade(score)
	fmt.Println(grade)

	// exe3
	loops()
}

// func sumofArray(a [5]int) int {
// 	println(a)
// 	sum := 0
// 	for i := range a {
// 		sum = sum + i
// 	}
// 	return sum
// }
