// Integer
// int 			== Platform dependent (Folowing Operating System)
// int8 		== -128 s.d 127
// int16 		== -32768 s.d 32767 (-215 to 215 -1)
// int32 		== -2147483648 s.d 2147483647 (231 to 231 -1)
// int64 		== -9223372036854775808 s.d 9223372036854775807 (-263 to 263 -1)

// Unsigned Integer
// uint			== Platform dependent (Folowing Operating System)
// uint8		== 0 s.d 255
// uint16		== 0 s.d 65535 (0 to 216 -1)
// uint32 	== 0 s.d 4294967295 (0 to 232 -1)
// uint64		== 0 s.d 18446744073709551615 (0 to 264 -1)

// Floating Point
// float32
// float64
// complex64
// complex128

// Alias
// byte = uint8
// rune = int32

package main

import "fmt"

func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga koma Lima = ", 3.5)
}
