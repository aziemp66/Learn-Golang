package main

import "fmt"

func main() {
	fmt.Println(sayHello("John", "Wick"))
}

func sayHello(firstName string, lastName string) string {
	return "Hello " + firstName + " " + lastName
}
