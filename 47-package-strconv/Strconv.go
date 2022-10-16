package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err1 := strconv.ParseBool("t")
	if err1 != nil {
		fmt.Println("Error : ", err1.Error())
		return
	}
	fmt.Println(boolean)

	number, err2 := strconv.ParseInt("1000000", 10, 64)
	if err2 != nil {
		fmt.Println("Error : ", err2.Error())
	}
	fmt.Println(number)

	value := strconv.FormatInt(100100, 16)
	fmt.Println(value)

	integer, err3 := strconv.Atoi("1000")
	if err3 != nil {
		fmt.Println("Error : ", err3.Error())
		return
	}
	fmt.Println(integer)

}
