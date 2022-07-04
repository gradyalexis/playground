package main

import "fmt"

func main() {
	sayHello("Grady", "Alexis")
}

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
