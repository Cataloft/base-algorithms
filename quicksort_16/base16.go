package main

import (
	"fmt"
	"math/rand"
)

func quickSort(slc []int) []int {
	if len(slc) < 2 {
		return slc
	}

	left := 0
	right := len(slc) - 1
	pivot := rand.Int() % len(slc)

	slc[pivot], slc[right] = slc[right], slc[pivot]

	for i, _ := range slc {
		if slc[i] < slc[right] {
			slc[left], slc[i] = slc[i], slc[left]
			left++
		}
	}

	slc[left], slc[right] = slc[right], slc[left]

	quickSort(slc[:left])
	quickSort(slc[left+1:])

	return slc
}

func generate(length int) []int {
	slc := make([]int, length)
	for i := 0; i < length; i++ {
		slc[i] = rand.Intn(length * 2)
	}
	fmt.Println(slc)
	return slc
}

func main() {
	var length int
	fmt.Scan(&length)
	slc := generate(length)
	fmt.Println(quickSort(slc))
}
