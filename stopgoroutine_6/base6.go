package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

// 1 способ: Отправка стоп-сигнала в канал
func waytostop1() {
	fmt.Println("This is from main function!")

	quitChan := make(chan bool)
	go func() {
		for {
			select {
			case <-quitChan:
				return
			default:
				// print a message every 3 seconds
				fmt.Println("This is test goroutine")
				time.Sleep(time.Second * 3)
			}
		}
	}()

	// sleep to print some message from the goroutine
	time.Sleep(time.Second * 10)

	// stop the goroutine
	quitChan <- true
	fmt.Println("Go rountine has stopped!")

	// test if the gorountine stopped or not
	time.Sleep(time.Second * 10)
	fmt.Println("This is the end of main func, goroutine no longer run!")

}

// 2 способ: Остановка горутины, закрыв канал
func waytostop2() {
	var wg sync.WaitGroup
	wg.Add(1)

	strChan := make(chan string)
	go func() {
		for {
			element, ok := <-strChan
			// check if channel is closed
			if !ok {
				println("Goroutines killed!")
				wg.Done()
				return
			}
			println(element)
		}
	}()
	strChan <- "this"
	strChan <- "is"
	strChan <- "a"
	strChan <- "message"
	close(strChan)

	// wait all goroutine to stop
	wg.Wait()
	// print the last message
	fmt.Println("This is the end of main func!")

}

// 3 способ: Остановить горутину, закрыв контекст
func waytostop3() {
	// a forever channel
	channel := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				channel <- struct{}{}
				return
			default:
				fmt.Println("This is from the gorountine")
			}

			// print a message every second
			time.Sleep(1 * time.Second)
		}
	}(ctx)

	// a go routine to close the above context
	go func() {
		time.Sleep(5 * time.Second)
		// close the context
		cancel()
	}()

	// run the gorountine forever
	<-channel
	fmt.Println("This is the end of the main func")

}

func main() {
	waytostop3()
}
