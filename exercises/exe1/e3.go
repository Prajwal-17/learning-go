package main

import "fmt"

func loops() {
	var arr [8]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}
	for i := range arr {
		fmt.Printf("index: %v & value: %v\n", i, arr[i])
	}
	count := 0
	for i := range arr {
		if arr[i] > 10 {
			count++
		}
	}
	fmt.Println("elements are greater that 10. count:", count)
	// fmt.Println("index multiple by 2", arr)

}
