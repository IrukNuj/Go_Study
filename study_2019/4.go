package main
import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age int
	phone string
}

// このメソッドを使ってHumanにfmt.Stringerを実装します。
func (h Human) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func main() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
