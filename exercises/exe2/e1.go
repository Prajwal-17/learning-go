// a simple in memory crud operation
package main

import (
	"fmt"
)

type User struct {
	id    int
	name  string
	age   int64
	email string
}

var allUsers = make(map[int]User)

func BasicCrud() {
	fmt.Println("CRUD")
	user1 := User{id: 1, name: "doe", age: 22, email: "doe@gmail.com"}

	user2 := User{id: 2, name: "john", age: 22, email: "john@gmail.com"}
	user3 := User{id: 3, name: "cena", age: 22, email: "cena@gmail.com"}

	// add
	user1.addUser()
	user2.addUser()
	user3.addUser()
	fmt.Println("allUsers map - ", allUsers)
	fmt.Println("result of key value 1 - ", allUsers[1])

	// delete
	delete(allUsers, 2)
	fmt.Println("After deleting - ", allUsers)

	// update
	// to update struct get value, modify, put if back
	m := allUsers[3] // read
	m.age = 50
	allUsers[3] = m
	fmt.Println("result after updating", allUsers)
	clear(allUsers)
	fmt.Println("result after clearing", allUsers)
}

func (u User) addUser() {
	allUsers[u.id] = u
}
