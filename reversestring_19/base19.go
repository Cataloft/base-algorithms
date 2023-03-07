package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

func reverseString(word string) string {
	r := []rune(word)

	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}

	return string(r)
}

func main() {
	var word string
	fmt.Scan(&word)

	fmt.Println(reverseString(word))
}
