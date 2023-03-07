package main

import "fmt"

/*Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.*/

// ждет пока в основной goroutine запишет число в 1 канал, считывает его возводя в квадрат и пишет во 2 канал
func writeSquare(ch1 chan int, ch2 chan int) {
	for v := range ch1 {
		ch2 <- v * v
	}
}

// ждет пока квадрат числа запишется в 2 канал и выводит в stdout
func printSquare(ch2 <-chan int) {
	for v := range ch2 {
		fmt.Println(v)
	}
}

// запускаем 2 goroutines и в основной в цикле пишем в 1 канал
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go writeSquare(ch1, ch2)
	go printSquare(ch2)

	for _, v := range x {
		ch1 <- v
	}
}
