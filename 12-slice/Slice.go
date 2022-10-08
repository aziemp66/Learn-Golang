package main

import "fmt"

func main() {
	myArray := []string{"John", "Wick", "Chaplin", "Doe", "Smith", "Dick", "Pete"}

	fmt.Println(myArray)

	mySlice := myArray[2:5]

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	mySlice2 := mySlice[2:]

	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))

	myArray[4] = "Will"
	fmt.Println(mySlice) // [Chaplin Doe Will] - mySlice is a reference to myArray

	mySlice[1] = "Doe Doe"
	fmt.Println(myArray) // [John Wick Chaplin Doe Doe Dick Pete] - myArray is a reference to mySlice

	// Append
	newArray := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	newSlice1 := newArray[5:]
	fmt.Println(newSlice1) // [Saturday Sunday]

	newSlice2 := append(newSlice1, "New Day")
	fmt.Println(newSlice2) // [Saturday Sunday New Day]

	newSlice2[2] = "Furday"
	fmt.Println(newSlice2) // [Saturday Sunday Furday]
	fmt.Println(newArray)  // [Monday Tuesday Wednesday Thursday Friday Saturday Sunday]

	//make
	newSlice3 := make([]string, 2, 5)
	newSlice3[0] = "Hello"
	newSlice3[1] = "World"
	fmt.Println(newSlice3) // [Hello World]
	fmt.Println(len(newSlice3))
	fmt.Println(cap(newSlice3))

	//copy
	newSlice4 := make([]string, len(newSlice3), cap(newSlice3))
	copy(newSlice4, newSlice3)
	fmt.Println(newSlice4) // [Hello World]

	//difference between slice and array
	//array
	myArray2 := [3]string{"John", "Wick", "Chaplin"}
	fmt.Println(myArray2) // [John Wick Chaplin]
	myArray3 := [...]string{"Januari", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(myArray3) // [Januari February March April May June July August September October November December]

	//slice
	var mySlice3 []string = []string{"John", "Wick", "Chaplin"}
	fmt.Println(mySlice3) // [John Wick Chaplin]
}
