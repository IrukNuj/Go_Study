package main

import "fmt"

func main() {
	a, b := 10, 10

	called(a, &b)

	fmt.Println(a)
	fmt.Println(b)
}

func called(a int, b *int) {
	a = a + 1
	*b = *b + 1
}
