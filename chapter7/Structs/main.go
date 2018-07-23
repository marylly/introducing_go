package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64{
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Circle struct {
	x float64
	y float64
	r float64
}

func circleArea(c *Circle) float64{
	return math.Pi * c.r*c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64{
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	// var c Circle
	// c := new(Circle)
	// c := Circle{x: 0, y: 0, r: 5}
	c := &Circle{0, 0, 5}
	fmt.Println(c)

	fmt.Println(c.x, c.y, c.r)
	c.x = 10
	c.y = 5

	c2 := Circle{0, 0, 5}
	fmt.Println(circleArea(&c2))

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}

