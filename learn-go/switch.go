package main

import "fmt"

func main() {
	name := "Grady"

	switch name {
	case "Grady":
		fmt.Println("Hello Grady")
	case "Alexis":
		fmt.Println("Who are you?")
	default:
		fmt.Println("Hmmm")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("Nama terlalu panjang")
	case length2 > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
