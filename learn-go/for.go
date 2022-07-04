package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Ini loop counter ke", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Ini loop counter ke", counter)
	}

	slice := []string{"Saya", "Sedang", "Makan"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	books := []string{"Belajar Golang", "Belajar Vuejs", "Belajar PHP"}

	// for _, book := range books (Gunakan _ jika index tidak dibutuhkan)
	for i, book := range books {
		fmt.Println("Index -", i, "=", book)
	}

	person := make(map[string]string)
	person["name"] = "Grady"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}
