package _struct

import (
	"fmt"
	"testing"
)

type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Println("Animal Eat")
}

type Dog struct {
	Animal
}

func (d *Dog) Eat() {
	fmt.Println("Dog Eat")
}

type Cat struct {
	Animal
}

func Test_Eat(t *testing.T) {
	d := Dog{}
	// will print: Dog Eat
	d.Eat()

	c := Cat{}
	// will print: Animal Eat
	c.Eat()
}
