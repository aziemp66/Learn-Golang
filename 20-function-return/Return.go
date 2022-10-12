package main

import "fmt"

func main() {
	firstName := "John"
	lastName := "Wick"

	result := sayHello(firstName, lastName)

	fmt.Println(result)
}

func sayHello(firstName string, lastName string) string {
	return "Hello " + firstName + " " + lastName
}
