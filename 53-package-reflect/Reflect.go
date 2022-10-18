package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
	Age  int `required:"false" max:"100"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	sample := Sample{"Eko", 30}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(1).Name)
	fmt.Println(sampleType.Field(1).Tag.Get("required"))
	fmt.Println(isValid(sample))
}
