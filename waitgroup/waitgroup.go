package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	wg := &sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			a[i] *= 3
			fmt.Println("goroutine", i)
			wg.Done()
		}(i) //invoca a funcao go anonima
	}

	wg.Wait()
	fmt.Println(a)
}
