package main

import "fmt"

func main() {
	firstName, middleName, lastName := getFullName()

	fmt.Println(firstName, middleName, lastName)
}

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Grady"
	middleName = "De"
	lastName = "Grace"

	// return firstName, middleName, lastName

	// cara singkat cukup return
	return
}
