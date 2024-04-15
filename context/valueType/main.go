package main

import (
	"context"
	"fmt"
)

type number int
type Person struct {
	Name string
}

func main() {
	ctx := context.WithValue(context.Background(), number(0), "value 1")
	ctx = context.WithValue(ctx, 0, "value 2")
	person := Person{Name: "saiteja"}
	ctx = context.WithValue(ctx, person, "person")
	handleRequest(ctx)

}

func handleRequest(ctx context.Context) {
	fmt.Println(ctx.Value(number(0)))
	fmt.Println(ctx.Value(int(0)))

	person := Person{Name: "saitejaa"}
	fmt.Println(ctx.Value(person))

}
