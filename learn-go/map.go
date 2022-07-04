package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string
	person := map[string]string{
		"name":    "Grady",
		"address": "Lembang",
	}

	// add map
	person["title"] = "Developer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])
	fmt.Println(len(person))

	var book map[string]string = make(map[string]string)
	book["title"] = "Learn Go"
	book["author"] = "Grady Alexis"
	book["release"] = "2020-12-12"
	book["ups"] = "Salah"

	delete(book, "ups")

	fmt.Println(book)
}
