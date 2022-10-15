package main

import "fmt"
import "belajar-golang-dasar/43-package-initialization/database"
import _ "os" //black identifier to allow package manaer feature

func main() {
	log := Database.GetDatabase()
	fmt.Println(log)
}
