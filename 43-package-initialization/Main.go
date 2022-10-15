package main

import "fmt"
import "belajar-golang-dasar/43-package-initialization/database"

func main() {
	log := Database.GetDatabase()
	fmt.Println(log)
}
