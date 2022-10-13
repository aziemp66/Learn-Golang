package main

import "fmt"

func main() {
	var goodBye func(string) string = getGoodBye

	fmt.Println(goodBye("John Wick"))
}

func getGoodBye(name string) string {
	return "Good bye " + name
}
