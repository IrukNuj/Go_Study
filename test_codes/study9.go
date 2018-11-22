package main

import (
	"fmt"
	"reflect"
)

var i, j int = 1, 2

func main()  {
	var C,python,java  = true, false , "no!"
	fmt.Println(i, j, C ,python,java)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(reflect.TypeOf(java)))
}