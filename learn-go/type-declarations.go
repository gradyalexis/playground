package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKTPGrady NoKTP = "129083128302198317"
	var marriedStatus Married = false

	fmt.Println("No. KTP:", noKTPGrady)
	fmt.Println("Status Menikah:", marriedStatus)
}
