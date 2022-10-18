package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Day(), "-", now.Month(), "-", now.Year(), ";", now.Hour(), ":", now.Minute(), ":", now.Second(), ":", now.Nanosecond())

	utc := time.Date(2022, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println(parse)
}
