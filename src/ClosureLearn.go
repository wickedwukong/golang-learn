package main

import "fmt"

func makeF() func(name string) (string, int) {
	x := 0
	return func(name string)(string, int) {
		x++
		return "Hello " + name, x
	}
}

func main() {
	f := makeF()

	fmt.Println(f("xuemin"))
	fmt.Println(f("xuemin"))
}