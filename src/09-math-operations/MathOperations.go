package main

import "fmt"

func main() {
	// 1. Math Operations
	var result = 10 + 10
	result = result - 10
	result = result * 10
	result = result / 10
	result = result % 10
	result++
	result--
	fmt.Println(result)

	// 2. Math Operations with Assignment
	var result2 = 10
	result2 += 10
	result2 -= 10
	result2 *= 10
	result2 /= 10
	result2 %= 10
	fmt.Println(result2)

	// 3. Math Operations with Decimals
	var result3 = 10.0
	result3 += 10.0
	result3 -= 10.0
	result3 *= 10.0
	result3 /= 10.0
	fmt.Println(result3)

	// 4. Math Operations with Decimals
	var result4 = 10.0
	result4 += 10.0
	result4 -= 10.0
}
