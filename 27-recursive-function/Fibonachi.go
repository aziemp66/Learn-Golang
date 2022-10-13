package main

func main() {
	println("Fibonachi: 5 =", fibonachi(10))
}

func fibonachi(value int) (count int) {
	if value == 1 || value == 2 {
		return 1
	}
	return fibonachi(value-1) + fibonachi(value-2)

}
