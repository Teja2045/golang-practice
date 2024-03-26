package main

import "fmt"

func main() {

outerLoops:

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		break outerLoops
	}

	fmt.Println("done")
}
