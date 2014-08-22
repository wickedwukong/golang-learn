package main

import (
	"fmt"
)

func main() {
	k := "hello world"
	fmt.Println(k)
	fmt.Printf("%T(%v)\n", "hello", "hello")
	i := -12.345
	f := int64(i)
	v := uint(f)

	fmt.Println(i, f, v)
	const world = "world"
	fmt.Println(world)
}