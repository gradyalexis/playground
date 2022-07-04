package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var user Customer
	user.Name = "Grady"
	user.Address = "Bandung"
	user.Age = 22

	user.sayHi("Alex")

	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Address)
	fmt.Println(user.Age)

	user2 := Customer{
		Name:    "Voler",
		Address: "Bandung",
		Age:     22,
	}
	fmt.Println(user2)

	user3 := Customer{"Dark Lord", "Bandung", 22}
	fmt.Println(user3)
}
