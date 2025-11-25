package main

// Error are values in go just like int, string
// In js or java errors are exceptions, when error occurs the ctrl flow jumps to catch block
// errors are handled immediately instead of bubbling up

//  error interface under the hood
//  type error interface {
//     Error() string
// }

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(4, 0)
	if err != nil {
		fmt.Println("Err occured", err)
		return
	}
	fmt.Println("divide value", result)
}
