package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var z [5]float64
	z[0] = 98
	z[1] = 93
	z[2] = 77
	z[3] = 82
	z[4] = 83

	var total float64 = 0
	for i := 0; i < len(z); i++ {
		total += z[i]
	}
	// fmt.Println(total / 5)
	// fmt.Println(total / len(z))
	fmt.Println(total / float64(len(z)))

	var total2 float64 = 0
	for _, value := range z {
		total2 += value
	}
	fmt.Println(total / float64(len(z)))
}