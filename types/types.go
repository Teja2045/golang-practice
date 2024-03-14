// types.go
package main

import (
	"fmt"
	"reflect"
)

type A1 struct{}
type A2 struct{}

type A interface {
	Print()
}

func (a A1) Print() {
	fmt.Println("A1 printing")
}

func (a A2) Print() {
	fmt.Println("A2 printing")
}

func PrintType(a A) {
	fmt.Println(reflect.TypeOf(a))
}

func PrintTypeWithoutReflect(a A) {
	switch v := a.(type) {
	case A1:
		fmt.Println("Type is A1")
	case A2:
		fmt.Println("Type is A2")
	case nil:
		fmt.Println("Type is nil")
	default:
		fmt.Printf("Type is unknown: %T\n", v)
	}
}
