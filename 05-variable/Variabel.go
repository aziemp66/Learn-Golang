package main

import "fmt"

func main() {
	var name string

	name = "Azie Melza Pratama"
	fmt.Println(name)

	name = "Fadhil Al-Fatih"
	fmt.Println(name)

	var money = 1000000
	fmt.Println(money)

	isMarried := false
	fmt.Println(isMarried)
	isMarried = true

	country := "Indonesia"
	fmt.Println(country)
	country = "Malaysia"
	fmt.Println(country)

	var (
		firstName = "Azie"
		lastName  = "Melza Pratama"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
