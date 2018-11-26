package main

import "fmt"

func main () {
	var pointer *int
	var n int = 100

	pointer = &n

	fmt.Println(&n) //nのアドレス
	fmt.Println(pointer) //pointer = &n
	fmt.Println(n) // n int = 100
	fmt.Println(*pointer)
	// *(間接参照演算子)
}