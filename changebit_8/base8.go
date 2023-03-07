package main

import (
	"fmt"
	"strconv"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
func convertInt(x int64) int64 {
	xHex := strconv.FormatInt(x, 2)
	var key int
	fmt.Println(len(xHex), xHex)
	fmt.Scan(&key)
	runes := []rune(xHex)
	fmt.Println(xHex)
	if runes[key-1] == '1' {
		runes[key-1] = '0'
	} else {
		runes[key-1] = '1'
	}
	fmt.Println(string(runes))
	res, _ := strconv.ParseInt(string(runes), 2, 10)
	return res

}

func main() {
	var x int64
	fmt.Scan(&x)
	fmt.Println(convertInt(x))
}
