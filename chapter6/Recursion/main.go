package main

import "fmt"

func main() {

	fmt.Println(fatorial(2))
}

func fatorial(x uint) uint{
	if x == 0 {
		return 1
	}

	return x * fatorial(x-1)
}