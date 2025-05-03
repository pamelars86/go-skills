package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		for i:= range ch1{
			defer fmt.Println(i)
		}
	}()
	
	ch1 <- 1
	ch1 <- 2
	fmt.Println("fim!")
	time.Sleep(2 * time.Second )
}