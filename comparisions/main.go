package main

import "fmt"

type Material struct {
	name     string
	strength int64
}
type Radio struct {
	volume   int64
	cost     int64
	material Material
}

type Speaker interface {
	makeSound()
}

func (rd *Radio) makeSound() {
	fmt.Println("radio is making noise..")
}

func (rd Radio) Compare1(rd2 Radio) {
	fmt.Println("comparision1: ", rd == rd2)
}

func (rd Radio) Compare2(rd2 Radio) {
	fmt.Println("comparision2: ", &rd == &rd2)
}

// deep copy
func (rd Radio) Compare3(rd2 *Radio) {
	fmt.Println("comparision3: ", &rd == rd2)
}

func (rd *Radio) Compare4(rd2 *Radio) {
	fmt.Println("comparision4: ", rd == rd2)
}

func main() {
	//var speaker Speaker

	material := Material{
		name:     "plastic",
		strength: 100,
	}
	rd := Radio{
		volume:   10,
		cost:     100,
		material: material,
	}

	//material.strength = 200

	rd2 := Radio{
		volume:   10,
		cost:     100,
		material: material,
	}

	fmt.Println(&rd == &rd2, rd == rd2, rd.material, rd2.material)

	rd.Compare1(rd2)
	rd.Compare2(rd2)
	rd.Compare3(&rd2)
	rd.Compare4(&rd2)
}
