// pointers & methods
package main

import "fmt"

type Item struct {
	name  string
	price float64
}

type Cart struct {
	items []Item
}

// c-> pointer to which type it is related  | i -> parameter, send actual object
func (c *Cart) addItem(i Item) {
	c.items = append(c.items, i)
}

func (c *Cart) getTotal() float64 {
	total := 0.0
	for _, item := range c.items {
		total += item.price
	}
	return total
}

func (c *Cart) clear() {
	c.items = []Item{}
}

func pointersMethods() {
	fmt.Println("Pointers & Methods")

	// type of Cart
	// in go a struct can contain a slice in it
	cart := Cart{}
	fmt.Println("Cart items", cart)

	cart.addItem(Item{name: "Bread", price: 234})
	cart.addItem(Item{name: "Milk", price: 40})
	cart.addItem(Item{name: "Tea", price: 10})

	fmt.Println("Adding items", cart)

	cart.clear()
	fmt.Println("After clearing", cart)

	// fmt.Printf("%+v\n", cartItems) // format printing
}
