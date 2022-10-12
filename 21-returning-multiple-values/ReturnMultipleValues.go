package main

func getFullName() (string, string) {
	return "John", "Doe"
}

func main() {
	firstName, lastName := getFullName()
	println(firstName, lastName)
}
