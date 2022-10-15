package main

func main() {
	println("Loop: 5! =", factiorialLoop(5))
	println("Recursive: 5! =", factiorialRecursive(5))
}

func factiorialLoop(value int) int {
	result := 1
	for value > 0 {
		result *= value
		value--
	}
	return result
}

func factiorialRecursive(value int) int {
	if value == 1 {
		return 1
	}
	return value * factiorialRecursive(value-1)
}
