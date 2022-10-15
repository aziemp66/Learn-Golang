package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address := Address{"Kayuagung", "Sumatera Selatan", "Indonesia"}

	changeCityAdress(address, "Indralaya")

	fmt.Println(address) // tidak berubah

	changeCityAddressWithPointers(&address, "Palembang")
	fmt.Println(address)

	// atau bisa juga dimasukkan variabel yg berisi pointer
	addressPointer := &address

	changeCityAddressWithPointers(addressPointer, "Linggau")

	fmt.Println(address)
}

func changeCityAdress(a Address, city string) {
	a.City = city
}

func changeCityAddressWithPointers(a *Address, city string) {
	a.City = city
}
