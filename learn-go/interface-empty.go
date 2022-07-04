package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	}

	if i == 2 {
		return 2
	}

	return "Ups"
}

func main() {
	var data interface{} = Ups(1)

	fmt.Println(data)
}
