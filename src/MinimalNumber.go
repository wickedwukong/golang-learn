package main

import "fmt"

func main() {
	x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17}

	fmt.Println(minimum(x))
}

func minimum(x []int) int {
	min := x[0]
	xs := x[1:]
	for _, v := range xs {
		if (min > v) {
			min = v
		}
	}
	return min
}