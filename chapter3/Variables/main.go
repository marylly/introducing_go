package main

import "fmt"

var b string = "Hello, World"

func main() {
	var x string = "Hello, World"
	fmt.Println(x)

	x = "Hello, World"
	fmt.Println(x)

	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	var z string = "hello"
	var y string = "world"
	fmt.Println( z == y )

	var w string = "hello"
	var t string = "hello"

	fmt.Println( w == t )

	v := "Hello, World"
	fmt.Println(v)

	r := 5
	fmt.Println(r)

	a := "Max"
	fmt.Println("My dog's name is", a)

	name := "Max"
	fmt.Println("My dog's name is", name)

	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)

	fmt.Println(b)
}

func f() {
	fmt.Println(x)
}