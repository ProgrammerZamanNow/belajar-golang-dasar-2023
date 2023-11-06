package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("Eko Kurniawan Khannedy")
}
