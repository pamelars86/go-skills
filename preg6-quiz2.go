package main

import "fmt"

func main() {
	var i interface{} = "hello world!"

	i = 1.3

	f := i.(float64)
	fmt.Println(f)
}


