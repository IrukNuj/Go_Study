package main

import (
	"fmt"
	"time"
)

func main () {
	hour := time.Now().Hour()

	switch hour {
	case 14:
		fmt.Println("hoge")
	default:
		fmt.Println("hugahuga")
	}

}