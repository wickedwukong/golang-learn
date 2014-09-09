package main

import "fmt"

func main() {
	fmt.Println("Please input your name")
	var name string
	fmt.Scanf("%s", &name)

	fmt.Println("hello", name)
}