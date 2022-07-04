package main

import "fmt"

func main() {
	const firstName string = "Grady"
	const lastName = "Alexis"
	const age int8 = 22

	fmt.Println("#Dengan banyak constant")
	fmt.Println(firstName, lastName)
	fmt.Println("Umur:", age)

	const (
		name    string = "Grady Alexis"
		address        = "Lembang"
	)

	fmt.Println("")
	fmt.Println("#Dengan satu constant")
	fmt.Println("Nama:", name)
	fmt.Println("Alamat:", address)
}
