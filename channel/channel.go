package main

import (
	"fmt"
	"sync"
)

func sendMessage(wg *sync.WaitGroup, msg string, ch chan string) {
	defer wg.Done()
	ch <- msg
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 2)
	ch2 := make(chan string, 2)

	wg.Add(2)
	go sendMessage(&wg, "Hello", ch)
	go sendMessage(&wg, "Bye", ch)

	wg.Add(2)
	go sendMessage(&wg, "Hello2", ch2)
	go sendMessage(&wg, "Bye2", ch2)

	go func() {
		wg.Wait()
		close(ch)
		close(ch2)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	for v := range ch2 {
		fmt.Println(v)
	}
}
