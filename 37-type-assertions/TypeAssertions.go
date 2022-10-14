package main

import "fmt"

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)
	//
	// resultInt := result.(int)
	// fmt.Println(resulInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String ", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown")
	}
}

func random() interface{} {
	return "OK"
}
