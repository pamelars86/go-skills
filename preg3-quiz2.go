package main

import "fmt"

func main() {
	var p *int
    var i interface{} = p

	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}



