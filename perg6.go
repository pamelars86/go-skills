package main

import "fmt"

func main() {
	x:= 100
	fmt.Println("Antes:", x)

	dobraValor(x)
	fmt.Println("Depois:", x)
}

func dobraValor(v int) {
	v = v * 2
}