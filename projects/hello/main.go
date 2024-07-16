package main

import (
	"fmt"
	"hello/contries"
	"log"
)

func div(a, b float32) (c float32) {
	log.Println("input parameters:", a, b)
	if b == 0 {
		log.Println("division by zero")
		return 0
	}

	c = a / b
	log.Printf("division by %g", c)
	return
}

func main() {
	fmt.Println("Hello World")

	fmt.Println("hello", contries.GetCountry("FR"))

	div(6.0, 3.0)
	div(6.0, 0.0)
}
