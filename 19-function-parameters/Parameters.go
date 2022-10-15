package main

import "fmt"

func main() {
	sayHello("John", "Wick")
	sayHello("Ethan", "Hunt")
}

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello " + firstName + " " + lastName)
}
