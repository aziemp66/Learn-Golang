package main

import (
	"belajar-golang-dasar/43-package-initialization/database"
	_ "belajar-golang-dasar/43-package-initialization/database" //Hanya Memanggil Init, ketika memanggil package yang sama init tidak dijalankan 2 kali
	"fmt"
)

func main() {
	log := database.GetDatabase()
	fmt.Println(log)
}
