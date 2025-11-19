package main

import "fmt"

var name string = "prajwal"
var phone int = 9844420817
var email string = "prajwal@gmail.com"

func variables() {
	var age int = 20
	var a, b int = 2, 4
	name = "rahul" // reassign name variable

	// shorthand for declaring and initializing
	clg_name := "bit"

	fmt.Println("variables")
	fmt.Println(name, phone, email, age)
	fmt.Println("numbers", a, b)
	fmt.Println("Clg Name - ", clg_name)
}
