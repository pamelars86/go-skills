package main

import (
	"fmt"
)

func main() {
	a := 10
	b:= &a
	*b = *b * 2
	fmt.Println(a, *b)
	
}