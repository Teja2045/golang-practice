// main.go
package main

// Import necessary packages here if needed

func main() {
	var a, b A
	a = A1{}
	b = A2{}
	a.Print()
	PrintType(a)
	PrintTypeWithoutReflect(a)
	PrintTypeWithoutReflect(b)
}
