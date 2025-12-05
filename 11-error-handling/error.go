package main

import (
	"errors"
	"fmt"
)

func divide2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by 0")
	}

	return a / b, nil // nil = "no error"
}

func errorHandling() {
	fmt.Println("error")
	result, err := divide2(4, 0)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result", result)
}
