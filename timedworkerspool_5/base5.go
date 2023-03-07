package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
//func read(out <-chan string) {
//	fmt.Println(<-out)
//}

func read(in <-chan string) {
	for val := range in {
		fmt.Println(val)
		time.Sleep(time.Second)
	}
}
func main() {
	in := make(chan string)
	var timed time.Duration
	_, err := fmt.Scan(&timed)
	if err != nil {
		log.Fatal(err)
	}
	go read(in)

	var slc []string

	i := 0
	timeout := time.After(timed * time.Second)
	for {
		slc = append(slc, strconv.FormatInt(int64(i+1), 10))
		select {
		default:
			in <- slc[i]
			i++
		case <-timeout:
			fmt.Println()
			return
		}

	}
	time.Sleep(time.Minute * 2)
}
