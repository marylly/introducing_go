package main

import "fmt"

var b string = "Hello, World"

func main() {
	var x string = "Hello, World"
	fmt.Println(x)

	fmt.Println(b)
}

func f() {
	fmt.Println(x)
}