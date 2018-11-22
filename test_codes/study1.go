package main

import "fmt"

func main()  {
	var num int //型は後ろに置く
	var pow int
	var result = 1

	num = 2
	pow = 4
	for  i:= 0; i < pow ; i++ { //波括弧は必須
		result *= num
	}
	fmt.Printf("%dの%d乗は%dです。¥n",num, pow, result)
}