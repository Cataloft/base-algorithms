package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.
type safeMap struct {
	sync.RWMutex
	nums map[int]int
}

func (s *safeMap) write(i int) {
	s.Lock()
	defer s.Unlock()
	s.nums[i] = i
}

func generateNums(numNums int) {
	wg := sync.WaitGroup{}

	numMap := &safeMap{
		nums: map[int]int{},
	}

	for i := 0; i < numNums; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			numMap.write(i)
		}(i)
	}
	fmt.Println(numMap.nums)
}

func main() {
	var numNums int
	fmt.Scan(&numNums)
	generateNums(numNums)
}
