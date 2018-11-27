package main

import "fmt"

type Vector struct {
	X int
	Y int //大文字じゃないと外部パッケージから読み取り不可
}

func main() {
	v := Vector{X: 5, Y: 2}

	fmt.Println(v)
	dayOfWeeks := [...]string{"月", "火", "水", "木", "金", "土", "日"}

	for arrayIndex, dayOfWeek := range dayOfWeeks {
		fmt.Printf("%d番目の曜日は%s曜日です。\n", arrayIndex+1, dayOfWeek)
	}

}
