package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("Aplikasi Error")
	}
}
