package main

import "fmt"

func main() {
	name := "Grady"

	if name == "Grady" {
		fmt.Println("Hello Grady")
	} else if name == "Gredy" {
		fmt.Println("Bangke")
	} else {
		fmt.Println("Anda siapa")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	benar := true

	if benar {
		fmt.Println("Benar")
	}
}
