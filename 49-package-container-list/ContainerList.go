package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Azie")
	data.PushBack("Melza")
	data.PushBack("Pratama")
	data.PushFront("Mr.")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	fmt.Println("\nFront to Back")
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("\nBack to Front")
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
