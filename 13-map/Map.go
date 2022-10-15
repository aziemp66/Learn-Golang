package main

import "fmt"

func main() {
	person := map[string]string{
		"name":        "Azie",
		"institution": "Unsri",
	}
	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["institution"])

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Azie"
	book["wrong"] = "ups"

	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))
}
