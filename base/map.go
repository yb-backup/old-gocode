package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var db = make(map[string]PersonInfo)

	// insert
	db["1"] = PersonInfo{"1", "yangbin", "Room 1,..."}
	db["2"] = PersonInfo{"2", "yangbin2", "Room 2,..."}

	person, ok := db["1"]
	if ok {
		fmt.Println("Found", person.Name, "with ID 1")
	} else {
		fmt.Println("Not Found person with ID 1")
	}
	person, ok = db["3"]
	if ok {
		fmt.Println("Found", person.Name, "with ID 3")
	} else {
		fmt.Println("Not Found person with ID 3")
	}
}
