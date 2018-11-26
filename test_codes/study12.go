package main

import (
	"fmt"
	"time"
)

func main () {
	now := time.Now()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	if hour >= 12 {
		fmt.Println("昼やで")
		fmt.Println(hour)
		fmt.Println(minute)
	}
	fmt.Println(now)
	fmt.Println(time.Now().Date())
}