package main
import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名フィールド
	school string
}

type Employee struct {
	Human //匿名フィールド
	company string
}

//Humanでmethodを定義
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//EmployeeのmethodでHumanのmethodを書き直す。
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}