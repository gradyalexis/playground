package main

import "fmt"

func main() {
	runApp(true)
	fmt.Println("Dark Lord")
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Application selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}
