package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer second()
	first()

	first()
	second()

	// f, _ := os.Open(filename)
	// defer f.Close()

	// panic("PANIC")
	// str := recover()
	// fmt.Println(str)

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}