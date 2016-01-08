package main

import "github.com/k0kubun/pp"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	pp.Print("Hi I am" + h.name + " " + h.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-2222"}, "MIT"}
	mark.SayHi()
	pp.Print(mark.name)
}
