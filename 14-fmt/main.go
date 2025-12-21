package main

import "fmt"

// https://pkg.go.dev/fmt

// %v -> default formats
// %T -> represents value
// %t -> boolean
// %b -> base 2
// %x -> base 16 with lowercase
// %X -> base 16 with uppercase
// %U -> Unicode
// %s -> string

// fmt.Println - is a part of fmt package, this writes to os.Stdout
// print - low level builtin function , this write to os.Stderr

func main() {
	name := "prajwal"
	n := int64(234)
	fmt.Printf("type is %T\n", name)
	fmt.Printf("Base 2 representation - %b\n", n)
	fmt.Printf("Base 8 representation - %o\n", n)
	fmt.Printf("Base 10 representation - %d\n", n)
	fmt.Printf("Base 16 representation - %x\n", n)
	fmt.Printf("Base 16 representation - %X\n", n)
	fmt.Printf("Unicode representation - %U\n", n)

	fmt.Printf("%s\n", name) // printf is for using format
	const name2, age = "Kim", 22
	fmt.Print(name2, " is ", age, " years old.\n")
	print("hello world\n")
}

// --- OUTPUT ---
/*
type is string
Base 2 representation - 11101010
Base 8 representation - 352
Base 10 representation - 234
Base 16 representation - ea
Base 16 representation - EA
Unicode representation - U+00EA
prajwal
Kim is 22 years old.
hello world
*/
