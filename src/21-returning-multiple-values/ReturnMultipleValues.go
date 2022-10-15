package main

func getFullName() (string, string, string) {
	return "John", "Doe", "Wick"
}

func main() {
	firstName, _, lastName := getFullName() // _ is a blank identifier (ignore value)
	println(firstName, lastName)
}
