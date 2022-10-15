package main

import "fmt"

func main() {
	runApp(true)
}

func endApp() {
	fmt.Println("Aplikasi selesai")
	message := recover()
	fmt.Println("Terjadi Error : ", message)

}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}
