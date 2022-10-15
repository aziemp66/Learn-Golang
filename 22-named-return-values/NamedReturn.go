package main

func main() {
	firstName, _, lastName := getFullName() // _ is a blank identifier (ignore value)
	println(firstName, lastName)
}
func getFullName() (firstName, middleName, lastName string) {
	firstName = "John"
	middleName = "Doe"
	lastName = "Wick"
	return
}
