package main

import (
	"fmt"
	"math/rand"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func randType() interface{} {
	r := rand.Intn(4)
	switch r {
	case 0:
		return 1
	case 1:
		return "cat"
	case 2:
		return true
	case 3:
		ch := make(chan interface{})
		return ch
	}
	return nil
}

func main() {
	var t interface{}
	t = randType()

	switch t := t.(type) {
	case int:
		fmt.Printf("Тип int: %d\n", t) // t has type bool
	case string:
		fmt.Printf("Тип string: %s\n", t) // t has type int
	case bool:
		fmt.Printf("Тип bool: %t\n", t) // t has type *bool
	case chan interface{}:
		fmt.Printf("Тип channel: %v\n", t) // t has type *int
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	}
}
