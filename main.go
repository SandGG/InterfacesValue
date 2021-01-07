package main

import (
	"fmt"
)

type animals interface {
	charac()
}

type attributes struct {
	name  string
	blood string
}

type birds struct {
	attributes
}

type reptile struct {
	attributes
}

func (a *attributes) charac() {
	fmt.Println("Name: " + a.name + ", Blood: " + a.blood)
	fmt.Printf("%T\n", a)
}

func printAnimal(a animals) {
	animals.charac(a)
}

func main() {
	printAnimal(&birds{attributes{name: "Eagle", blood: "hot"}})
	printAnimal(&reptile{attributes{name: "Cocodrile", blood: "cold"}})
}
