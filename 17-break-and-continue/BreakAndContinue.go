package main

func main() {
	for i := 1; i <= 50; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		println("Angka", i)
	}
}
