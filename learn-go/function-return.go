package main

import "fmt"

func main() {
	result := getHello("Grady")

	fmt.Println(result)
}

func getHello(name string) string {
	if name != "" {
		return "Hello " + name
	}

	return "Hello bro"
}
