package main

import "fmt"

func main() {
	xs := []float64{93,93,77,82,83}

	fmt.Println(average(xs))

	x := 5
	f(x)

	fmt.Println(f1())

	z, y := f3()
	fmt.Println(z)
	fmt.Println(y)
}

func average( xs []float64) (total float64) {
	for _, v := range xs {
		total += v
	}
	total = total / float64(len(xs))
	return
}

func f(x int) {
	fmt.Println(x)
}

func f1() int{
	return f2()
}

func f2() int{
	return 1
}

func f3() (int, int) {
	return 5, 6
}