package main

import "fmt"

func send(value string, ch chan string, done chan bool) {
	ch <- value
	done <- true
}

func main() {
	c := make(chan string, 2)
	done := make(chan bool, 2)

	go send("hello", c, done)
	go send("world", c, done)

	go func() {
		<-done
		<-done
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
