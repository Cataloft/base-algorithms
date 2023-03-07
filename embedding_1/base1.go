package main

import "fmt"

//1.Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	name string
	age  int
}

func (h *Human) PrintName() {
	fmt.Printf("Name is %s\n", h.name)
}

func (h *Human) PrintAge() {
	fmt.Printf("Age is %d\n", h.age)
}

type Action struct {
	Human
}

func main() {
	var x Action
	x.name = "Slava"
	x.age = 23
	x.PrintName()
	x.PrintAge()
}
