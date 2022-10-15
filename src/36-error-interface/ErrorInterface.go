package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Dengan Nol")
	}
	return nilai / pembagi, nil
}

func main() {
	// contohError := errors.New("Contoh Error")
	// fmt.Println(contohError.Error())

	result1, result1Error := pembagian(10, 5)
	result2, result2Error := pembagian(30, 0)

	if result1Error != nil {
		fmt.Println(result1Error.Error())
	} else {
		fmt.Println(result1)
	}

	if result2Error != nil {
		fmt.Println(result2Error.Error())
	} else {
		fmt.Println(result2)
	}
}
