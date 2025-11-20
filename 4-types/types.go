package main

import "fmt"

var isActive bool = false
var name string = "john"

// integer
// - signed => int, int8, int16, int32, int64
// - unsigned => uint, uint8, uint16, uint32, uint64
var age int = 20
var b byte = 255 // store bytes

func main() {
	fmt.Println(isActive)
	fmt.Printf("Type of %v is: %T\n", isActive, isActive) // print typeof variable

	// -------

	s := "hi"
	// slice of bytes, this is not an array , array is fixed size but slice is dynamic
	ba := []byte(s)
	fmt.Println(ba) // prints ascii byte values

	// -------

	// represents unicode value
	var r rune = 'a'
	fmt.Println("Unicode value", r)

	// -------

	// float32 - 32bits | float64(default) - 64bits
	var price float64 = 283.55
	fmt.Println("price", price)

	// -------

	// complex64 - 64bits | complex128 - 128bits
	// complex => float<bits> real + float<bits> imaginary
	var c complex64 = 2 + 3i
	fmt.Println("value of complex", c)
}
