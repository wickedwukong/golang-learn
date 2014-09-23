package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(value int) {
	for i := 0; i < 10; i++ {
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
		fmt.Println(value, ":", i)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}