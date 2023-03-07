package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sort"
)

//Реализовать бинарный поиск встроенными методами языка.

func binarySearch(slc []int, x int) (int, error) {
	left := 0
	right := len(slc) - 1
	mid := 0
	for left <= right {
		mid = left + (right-left)/2
		if slc[mid] == x {
			return mid, nil
		} else if slc[mid] < x {
			left = mid + 1
		} else if slc[mid] > x {
			right = mid - 1
		}
	}
	return 0, errors.New("Искомый элемент не найден")
}

func generate(length, x int) []int {
	slc := make([]int, length)
	for i := 0; i < length; i++ {
		if rand.Intn(length-1) == i {
			slc[i] = x
			continue
		}
		slc[i] = rand.Intn(length * 2)
	}
	sort.Ints(slc)
	fmt.Println(slc)
	return slc
}

func main() {
	var length int
	var x int
	fmt.Scan(&x, &length)
	//fmt.Println(rand.Intn(length - 1))
	slc := generate(length, x)
	res, err := binarySearch(slc, x)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID искомового элемента: %d", res)
}
