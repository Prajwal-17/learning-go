package main

import "fmt"

// a struct is a collection of fields
// a user defined type that allows to group fields under single name

type person struct {
	name string
	age  int
}

type Vertex struct {
	X int
	Y int
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r *rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {
	fmt.Println(person{"joe", 21})
	fmt.Println(Vertex{234, 352})
	p2 := person{"john", 34}
	fmt.Println(p2.name)
	fmt.Println(p2.age)
	//-------
	r := rect{34, 23}
	fmt.Println("area", r.area())
	fmt.Println("perimeter", r.perim())
}
