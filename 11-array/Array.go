package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "John"
	names[1] = "Wick"
	names[2] = "Chaplin"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var names2 = [3]int{
		20,
		30,
		40,
	}

	fmt.Println(names2)
	fmt.Println(len(names2))
}
