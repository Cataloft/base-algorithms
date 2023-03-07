package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func ownSleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	var sleepTime time.Duration
	fmt.Println("Введите продолжительность сна в секундах: ")
	fmt.Scan(&sleepTime)

	fmt.Printf("Спим %d секунд\n", sleepTime)
	ownSleep(sleepTime * time.Second)
	fmt.Println("Просыпаемся")
}
