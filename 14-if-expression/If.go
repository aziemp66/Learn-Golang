package main

import "fmt"

func main() {
	name := "Pratama"

	if name == "Azie" {
		fmt.Println("Hello Azie")
	} else if name == "Melza" {
		fmt.Println("Hello Melza")
	} else {
		fmt.Println("Hello", name, "! Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
