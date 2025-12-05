package main

import (
	"errors"
	"fmt"
)

func ValidateAge(age int) (bool, error) {
	if age > 150 {
		return false, errors.New("Age is realistically too high")
	}

	return true, nil

}

func main() {
	result, err := ValidateAge(345)

	if err != nil {
		fmt.Println("An Error occured", err)
		return
	}
	fmt.Println(result)
}
