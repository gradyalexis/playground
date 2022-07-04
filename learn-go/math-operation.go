package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println("10 + 10 =", result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println("10 * 10 =", c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println("i += 10 =", i)

	i++
	fmt.Println("i++ =", i)

	var positive = +100 // 100
	var negative = -100

	fmt.Println("Positive =", positive)
	fmt.Println("Negative =", negative)
}
