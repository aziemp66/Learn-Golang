package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) string {
	return ("Hello " + name + " My Name is " + customer.Name)
}

func main() {
	Azie := Customer{Name: "Azie"}

	result := Azie.sayHello("Budi")

	fmt.Println(result)
}
