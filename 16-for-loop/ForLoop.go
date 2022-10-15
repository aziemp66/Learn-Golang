package main

import "fmt"

func main() {
	//for i
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke", i)
	}
	//for ranges
	slice := []string{"John", "Wick", "Doe", "Jacob", "Smith"}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}
}
