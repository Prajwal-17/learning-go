package main

import (
	"fmt"
)

// a interface is a collection of method signatures
type Geometry interface {
	perim() int
}

type rect struct {
	width, height int
}

func (r *rect) perim() int {
	return 2*r.height + 2*r.width
}

func measure(g Geometry) {
	fmt.Println("perimeter", g.perim())
}

func main() {
	x := rect{width: 3, height: 34}
	measure(&x)
}
