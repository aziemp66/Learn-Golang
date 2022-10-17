package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{"Azie", 300},
		{"Melza", 140},
		{"Pratama", 500},
		{"Vanhautten", 100},
	}

	userSlice := UserSlice(users)

	fmt.Println(userSlice)
	sort.Sort(userSlice)

	fmt.Println(userSlice)
}
