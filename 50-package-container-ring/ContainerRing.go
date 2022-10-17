package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	// data.Value = "Azie"

	// var data2 = data.Next()
	// data2.Value = "Melza"

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data : " + strconv.Itoa(i)
		data = data.Next()
	}

	fmt.Println(*data)

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
