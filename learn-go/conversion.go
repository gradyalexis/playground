package main

import "fmt"

func main() {
	// Integer Conversion
	// hati-hati saat melakukan konversi tipe data
	// apabila tipe data yang dikonversi terlalu besar maka akan terjadi overflow
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// String Conversion
	var name = "Grady Alexis"
	var e byte = name[0]           // Mengambil karakter pertama
	var eString string = string(e) // Konversi byte karakter pertama ke string

	fmt.Println(name)
	fmt.Println(eString)
}
