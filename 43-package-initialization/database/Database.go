package database

import "fmt"

var connection string

func init() {
	fmt.Println("Init Dipanggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
