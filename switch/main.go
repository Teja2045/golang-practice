package main

import "fmt"

func main() {
	d := 10

	switch d {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 10:
		fmt.Println("ten")
		fallthrough
	case 100:
		fmt.Println("hundred")
	default:
		fmt.Println("default")
	}

	fmt.Println("done")
}
