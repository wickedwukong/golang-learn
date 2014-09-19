package main

import "fmt"

type Person struct {
	Name string
}

func (person *Person) Talk() {
	fmt.Println("Hi, my name is", person.Name)
}

func main() {
	person := Person{Name: "Xuemin"}

	person.Talk()
}

