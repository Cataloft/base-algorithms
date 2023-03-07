package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func addition(x *big.Int, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

func subtraction(x *big.Int, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

func division(x *big.Int, y *big.Int) *big.Float {
	xF := new(big.Float).SetInt(x)
	yF := new(big.Float).SetInt(y)
	return new(big.Float).Quo(xF, yF)
}

func multiplication(x *big.Int, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}
func main() {
	var a, b string
	fmt.Scan(&a, &b)
	x, _ := new(big.Int).SetString(a, 10)
	y, _ := new(big.Int).SetString(b, 10)
	fmt.Println("Операция сложения данных чисел: ", addition(x, y))
	fmt.Println("Операция вычитания данных чисел: ", subtraction(x, y))
	fmt.Println("Операция деления данных чисел: ", division(x, y))
	fmt.Println("Операция умножения данных чисел: ", multiplication(x, y))
}
