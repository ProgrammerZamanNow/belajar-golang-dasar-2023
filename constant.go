package main

import "fmt"

func main() {
	const (
		firstName string = "Eko"
		lastName         = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Budi"
	// lastName = "Joko"
}
