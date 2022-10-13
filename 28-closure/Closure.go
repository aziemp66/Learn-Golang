package main

func main() {
	name := "Azie"
	counter := 0

	increment := func() {
		name := "Aziz" // if you do not declare a new variable, the value of the variable will be changed
		println(name, "incrementing")
		counter++

	}

	increment()
	increment()

	println(name, "counter =", counter)
}
