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

func primeNumbers() []int {
	return []int{2, 3, 5, 7, 11, 13}
}
