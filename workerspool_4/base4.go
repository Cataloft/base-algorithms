package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func worker(id int, in chan string, results chan string) {
	for val := range in {
		fmt.Println("worker", id, "started job")
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job and read: ", val)
		results <- val
	}
}

func genAndWrite(in chan string, results chan string, buf int) {
	var slc []string
	i := 0
	k := 0

	go interruptProgramm(in, results)

	for {
		slc = append(slc, strconv.FormatInt(int64(i+1), 10))
		if k == buf {
			fmt.Println("~~~Считывание с канала~~~")
			for a := 1; a <= buf; a++ {
				<-results
			}

			in <- slc[i]
			k = 1
			i++
			continue
		}
		in <- slc[i]
		k++
		i++
	}
}

func interruptProgramm(in, res chan string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for sig := range c {
		fmt.Println("exiting programm by", sig)
		os.Exit(1)
	}
}

func main() {
	var workers int
	_, err := fmt.Scan(&workers)
	if err != nil {
		log.Fatal(err)
	}

	buf := 5
	in := make(chan string, buf)
	results := make(chan string, buf)
	for w := 1; w <= workers; w++ {
		go worker(w, in, results)
	}

	genAndWrite(in, results, buf)
}
