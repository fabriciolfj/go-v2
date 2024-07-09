package main

import (
	"fmt"
	"sync"
)

func sendMsg(value string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- value
}

func main() {
	var messages = []string{"hello", "world", "from", "go"}

	ch := make(chan string, len(messages))
	var wg sync.WaitGroup

	wg.Add(len(messages))

	for _, message := range messages {
		go sendMsg(message, ch, &wg)
	}

	go func() {
		wg.Wait() // Espera todas as goroutines terminarem
		close(ch) // Então fecha o canal
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
