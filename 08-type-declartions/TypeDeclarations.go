package main

import "fmt"

func main() {
	type noKTP string
	type married bool

	var ekoKTP noKTP = "1234567890"
	fmt.Println(ekoKTP)
	var marriedStatus married = true
	fmt.Println(marriedStatus)
}
