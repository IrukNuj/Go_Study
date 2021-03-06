package main

import "fmt"

func later() func(string) string {
	var store string // クロージャ内から参照されている変数
	return func (next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := later()

	fmt.Println(f("Golang"))
	fmt.Println(f("is"))
	fmt.Println(f("Awesome"))
	fmt.Println(f(""))
}