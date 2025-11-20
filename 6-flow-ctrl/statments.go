package main

import "fmt"

// there is not ternary operator in go lang

func statments(x float64) {
	if x < 0 {
		fmt.Println("the number is less than 0")
	} else {
		fmt.Println("the number is greater than 0")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num := 23; num < 10 {
		fmt.Println(num, "is greater that 10 ")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
