package main

type CanTalk interface {
	Talk() string
}

type Animal struct {
	Species string
	Voice   string
}

func (animal Animal) Talk() string {
	// return animal.Voice
	return animal.Voice
}

func main() {
	var animal Animal
	animal.Species = "Cat"
	animal.Voice = "Meow"

	animal.Talk()
}
