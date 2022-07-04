package main

import "fmt"

func main() {
	var months = [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[10:]

	// using make() to create slice
	var slice3 = append(slice2, "Grady")
	fmt.Println(slice3)
	slice3[1] = "Alexis"
	fmt.Println(slice3)
}
