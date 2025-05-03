package main

import "fmt"

func main() {
	ch := make(chan int, 1) // buffer de 1

	go func() {
		ch <- 1
		ch <- 2
	}()

	fmt.Println(<-ch) // Imprime 1
	fmt.Println(<-ch) // Imprime 2
}
