package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.

func uniqueSymbols(symbols string) bool {
	var contains bool
	sb := ""
	for i, v := range symbols {
		contains = strings.Contains(sb+symbols[i+1:], string(v))
		sb += string(v)
		if contains {
			return false
		}
	}
	return true
}

func main() {
	var symbols string
	fmt.Scan(&symbols)
	fmt.Println(uniqueSymbols(symbols))
}
