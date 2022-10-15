package main

import "fmt"

type Man struct {
	Name string
}

func (m *Man) Married() {
	m.Name = "Mr." + m.Name
	fmt.Println(m.Name)
}
func main() {
	azie := Man{"Azie"}

	azie.Married()

	fmt.Println(azie)
}
