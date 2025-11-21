package main

import "fmt"

// if both params are same type then -> x, y int
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("result is", add(2, 3))
	sum, sub := calculator_1(5, 4)
	fmt.Println(sum, sub)
	sum_2, sub_2 := calculator_2(5, 32)
	fmt.Println(sum_2, sub_2)
}

// multiple return
func calculator_1(a int, b int) (int, int) {
	return a + b, a - b
}

func calculator_2(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

// --------------
// Functions as argument
// package main

// import "fmt"

// func sum(a int, b int) int {
// 	return a + b
// }

// func sub(a int, b int) int {
// 	return a - b
// }

// func calculator(a int, b int, fn func(int, int) int) int {
// 	return fn(a, b)
// }

// func main() {
// 	ans := calculator(1, 2, sum)

// 	fmt.Println(ans)
// }

// -----------
// returning functions

// package main

// import "fmt"

// func multiplier(factor int) func(int) int {
// 	return func(a int) int {
// 		return a * factor
// 	}
// }

// func main() {
// 	double := multiplier(3)
// 	fmt.Println(double(3))
// }
