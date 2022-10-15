package main

import "fmt"

func main() {
	log("Hello")
	log(4)
}

func log(any interface{}) interface{} {
	fmt.Println(any)
	return any
}
