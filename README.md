# Goroutines
```
	for i := 0; i < 5; i++ {
		go func(i int) {
			a[i] *= 3
			fmt.Println("goroutine", i)
			wg.Done()
		}(i) //invoca a funcao go anonima
	}
```