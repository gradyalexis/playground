package main

import "fmt"

func main() {
	// var name string = "Grady Alexis"
	// var name = "Grady Alexis"
	var name string

	name = "Grady Alexis"
	fmt.Println(name)

	name = "Grady"
	fmt.Println(name)

	var age uint8 = 22
	fmt.Println("Umur:", age)

	address := "Lembang"
	fmt.Println("Alamat:", address)

	var (
		firstName = "Grady"
		lastName  = "Alexis"
	)
	fmt.Println(firstName, lastName)
}
