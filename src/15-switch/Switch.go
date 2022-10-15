package main

import "fmt"

func main() {
	name := "Pratama"

	switch name {
	case "Azie":
		fmt.Println("Hello Azie")
	case "Melza":
		fmt.Println("Hello Melza")
	default:
		fmt.Println("Hello", name, "! Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}
}
