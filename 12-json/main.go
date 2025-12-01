package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Linus", "Hi there", 23423}
	b, err := json.Marshal(m)
	// Marshal outputs in bytes
	fmt.Println(b)
	fmt.Println("Json -", string(b), err)

	jsonInput := []byte(`{"Name":"Bob", "Body":"Checking in", "Time":1704067200, "Food":"Pickle"}`)
	var d Message
	unmarshalErr := json.Unmarshal(jsonInput, &d)
	fmt.Println(unmarshalErr)
	fmt.Println("Original JSON Input:", string(jsonInput))
}
