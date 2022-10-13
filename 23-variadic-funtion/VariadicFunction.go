package main

func main() {
	myTotal := sumAll(10, 30, 50, 112, 300)
	println(myTotal)

	//slice

	slice := []int{10, 30, 50, 112, 300}
	myTotal = sumAll(slice...)
	println(myTotal)

	//array

	myArray := [...]int{15, 30, 45, 60, 75, 90, 105, 120, 135, 150}
	myTotal = sumAll(myArray[:]...)
	println(myTotal)
}

func sumAll(numbers ...int) (total int) {
	total = 0

	for _, number := range numbers {
		total += number
	}

	return
}
