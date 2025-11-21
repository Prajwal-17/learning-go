package main

import "fmt"

// maps are Go built in data types that are key value pairs(hash-based)
// they are unordered collection

func maps() {
	m := make(map[string]int) // keys string & int values
	m["k1"] = 230
	m["k2"] = 324
	fmt.Println("map", m)

	names := map[string]int{
		"John": 25,
		"Bob":  24,
	}

	fmt.Println("map", names)

	delete(m, "k1")
	fmt.Println("delete", m)

	clear(m)
	fmt.Println("clear", m)
}

// map map[k1:230 k2:324]
// map map[Bob:24 John:25]
// delete map[k2:324]
// clear map[]
