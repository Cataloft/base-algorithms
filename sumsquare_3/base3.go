package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func squareNum(num int) int {
	return num * num
}

func handler(nums []int) int {
	sq := 0
	wg := sync.WaitGroup{}
	wg.Add(len(nums))
	for _, num := range nums {
		go func(val int) {
			defer wg.Done()
			sq += squareNum(val)
		}(num)
	}
	wg.Wait()
	return sq
}

func main() {
	var nums []int
	nums = append(nums, 2, 4, 6, 8, 10)
	fmt.Println(handler(nums))
}
