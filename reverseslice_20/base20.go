package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func reverseStrings(words string) string {
	slc := strings.Split(words, " ")

	for i := 0; i < len(slc)/2; i++ {
		slc[i], slc[len(slc)-1-i] = slc[len(slc)-1-i], slc[i]
	}

	rev := strings.Join(slc, " ")

	return rev
}

func main() {
	words := "snow dog sun"
	fmt.Println(words)

	fmt.Println(reverseStrings(words))
}
