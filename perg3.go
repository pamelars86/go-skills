package main

import "fmt"

type Gopher struct {
	Name string
	Winner bool
}

func main() {
	gophers := []Gopher{
		{"João", false},
		{"Maria", false},
		{"José", false},
	}
	
	for _, g := range gophers {
		g.Winner = true
	}
	
	var winners int
	
	for _, g := range gophers {
		if g.Winner {
			winners++
		}
	}
	fmt.Println(winners)
}
