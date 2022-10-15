package main

import "fmt"

func main() {
	const firstName string = "Azie"
	const lastName = "Melza Pratama"

	//error
	// firstName = "Fadhil"
	// lastName = "Al-Fatih"

	const (
		firstName2 string = "Azie"
		lastName2         = "Melza Pratama"
	)

	fmt.Println(firstName2)
	fmt.Println(lastName2)
}
