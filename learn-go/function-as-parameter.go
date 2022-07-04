package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	nameFiltered := filter(name)

	return "Hello " + nameFiltered
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name
}

func main() {
	fmt.Println(sayHelloWithFilter("Grady", spamFilter))
	fmt.Println(sayHelloWithFilter("Anjing", spamFilter))
}
