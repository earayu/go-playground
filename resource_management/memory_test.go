package resource_management

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
}

func createPerson() Person {
	p := Person{
		Name: "Curry",
	}
	fmt.Printf("The address of p is %p\n", &p)
	return p
}

func Test_VariableCreateInFuctionThenReturnAsInstance(t *testing.T) {
	p := createPerson()
	fmt.Printf("The address of p is %p\n", &p)
}

func createPersonPointer() *Person {
	p := Person{
		Name: "Curry",
	}
	fmt.Printf("The address of p is %p\n", &p)
	return &p
}

func Test_VariableCreateInFuctionThenReturnAsPointer(t *testing.T) {
	p := createPersonPointer()
	fmt.Printf("The address of p is %p\n", p)
}
