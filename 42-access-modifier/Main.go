package main

import (
	"belajar-golang-dasar/42-access-modifier/access"
	"fmt"
)

func main() {
	fmt.Println(access.SayHello("Azie"))
	// fmt.Println(access.sayGoodbye("Azie"))

	fmt.Println("Application Version : ", access.ApplicationVersion)
	// fmt.Println("Api Keys : ", access.apiKeys)
}
