package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	//pass by value

	address1 := Address{"Kayuagung", "Sumatera Selatan", "Indonesia"}
	address2 := address1

	address2.City = "Indralaya"

	// Data tidak sama
	// Karena data di address1 tidak berubah
	fmt.Println(address1)
	fmt.Println(address2)

	//pass by reference
	//& Pointer
	var address3 Address = Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}
	var address4 *Address = &address3

	address4.City = "Jakarta Selatan"

	fmt.Println(address3)
	fmt.Println(address4)

	//* Pointer
	address5 := Address{"California", "California", "USA"}
	address6 := &address5

	address6.City = "Cali"

	// address6 = &Address{"New Jersey", "Boston", "USA"}
	//
	// fmt.Println(address5) //address5 tidak berubah
	// fmt.Println(address6)

	*address6 = Address{"Batu Raja", "Sumatera Selaatan", "Indonesia"}

	fmt.Println(address5)
	fmt.Println(address6)

	//Function New
	address7 := new(Address)
	address8 := address7

	address8.Country = "USA"

	fmt.Println(address7)
	fmt.Println(address8)
}
