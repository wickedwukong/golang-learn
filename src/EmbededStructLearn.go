package main

import "fmt"

type Person struct {
	Name string
}

func (person *Person) Talk() {
	fmt.Println("Hi, my name is", person.Name)
}

type Android struct {
	Person Person
	Model string
}

type Droid struct {
	Person
	Model string
}

func main() {
	person := Person{Name: "Xuemin"}

	person.Talk()

	seeTHREEPIO := Android{Person: Person{Name: "SEE-THREEPIO"}, Model: "C-3PO"}

	seeTHREEPIO.Person.Talk()

	r2D2 := Droid{Person: Person{Name: "R2-D2"}, Model: "R2-D2"}
	r2D2.Talk()
}

