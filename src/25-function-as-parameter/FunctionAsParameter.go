package main

import (
	"strings"
)

type Filter func(string) string

func main() {
	result := sayHelloWithFilter("kiki", wordsFilter)
	println(result)
}

func wordsFilter(name string) string {
	if name == "Bajingan" {
		return "World"
	} else {
		return name
	}
}

func upperFilter(name string) string {
	return strings.Title(name)
}

func sayHelloWithFilter(name string, filter Filter) string {
	nameFiltered := filter(name)
	caseFiltered := upperFilter(nameFiltered)
	return "Hallo " + caseFiltered
}
