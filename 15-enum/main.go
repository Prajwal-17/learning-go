package main

import "fmt"

// there is no built in keyword as "enum" in go

// iota is a identifier in go that generates constant values automatically (0,1,2)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
)

type TrendOption string

const (
	INCREASE TrendOption = "increase"
	DECREASE TrendOption = "decrease"
	NOCHANGE TrendOption = "nochange"
)

func main() {
	fmt.Println(Sunday, Monday, Tuesday)
	fmt.Println(INCREASE, DECREASE, NOCHANGE)
}
