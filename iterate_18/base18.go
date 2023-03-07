package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	sync.RWMutex
	num int
}

func (c *Counter) increment() {
	c.Lock()
	defer c.Unlock()

	c.num += 1
}

func someJob(workers int, counter *Counter, end chan struct{}) {
	wg := sync.WaitGroup{}

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func(i int, counter *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("Какая-то работа выполяется воркером: ", i+1)
			counter.increment()
		}(i, counter, &wg)
	}

	wg.Wait()
	end <- struct{}{}
	close(end)
}

func main() {
	var workers int
	fmt.Scan(&workers)

	counter := &Counter{
		num: 0,
	}

	end := make(chan struct{})

	go someJob(workers, counter, end)

	_, ok := <-end
	if ok {
		fmt.Printf("Счетчик = %d", counter.num)
	}

}
