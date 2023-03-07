package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func squareNum(num int) int {
	return num * num
}

func handler(nums []int) []int {
	sq := make([]int, len(nums))
	wg := sync.WaitGroup{}
	wg.Add(len(nums))
	for i, num := range nums {
		go func(idx, val int) {
			defer wg.Done()
			sq[idx] = squareNum(val)
		}(i, num)
	}
	wg.Wait()
	return sq
}

func main() {
	var nums []int
	nums = append(nums, 2, 4, 6, 8, 10)
	fmt.Println(handler(nums))
}
