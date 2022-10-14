package main

import "fmt"

func main() {
	myMap := NewMap("Azie")

	if myMap == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(myMap)
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}
