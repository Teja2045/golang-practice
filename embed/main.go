package main

import "fmt"

type Head struct {
	size int16
}

func (head *Head) changeSize(size int16) {
	head.size = size
}

type Animal struct {
	Name string
	Age  int32
	Head
}

// direct embedding, we can have access to the methods of animals directly
type ZooEmbed struct {
	Animal
}

// with reference, we loose the direct embedding advantages
type Zoo struct {
	Animal Animal
}

func (animal *Animal) ChangeName(name string) {
	animal.Name = name
}

func main() {
	var zoo *Zoo
	var zooEmbed *ZooEmbed
	zoo = &Zoo{
		Animal{
			Name: "hemanth",
			Age:  10,
		},
	}

	zooEmbed = &ZooEmbed{
		Animal{
			Name: "hemanth 1",
			Age:  100,
		},
	}

	var head *Head
	var head2 *Head

	// head = &zoo.Head can't do this
	head2 = &zoo.Animal.Head

	// can do this since this is embedded thing
	head = &zooEmbed.Head

	fmt.Println(zoo)
	// zoo.ChangeName("teja")
	// zoo.changeSize(10) can't do

	zooEmbed.ChangeName("teja")
	zooEmbed.changeSize(10)

	fmt.Println(zoo, head, head2, zooEmbed)
}
