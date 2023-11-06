package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var eko Customer
	fmt.Println(eko)

	eko.Name = "Eko Kurniawan"
	eko.Address = "Indonesia"
	eko.Age = 30
	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	budi.sayHello("Agus")
	eko.sayHello("Agus")
	joko.sayHello("Agus")
}
