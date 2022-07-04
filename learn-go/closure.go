package main

import "fmt"

func main() {
	name := "Grady"
	counter := 0

	increment := func() {
		name = "Dark Lord"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
