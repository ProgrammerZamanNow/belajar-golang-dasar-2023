package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
