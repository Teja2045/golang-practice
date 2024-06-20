package main

import (
	"fmt"
	"sort"
)

type Integers struct {
	data []uint64
}

func (ints Integers) Len() int {
	return len(ints.data)
}

func (ints Integers) Less(i, j int) bool {
	return ints.data[i] < ints.data[j]
}

func (ints Integers) Swap(i, j int) {
	ints.data[i], ints.data[j] = ints.data[j], ints.data[i]
}

func main() {

	var _ sort.Interface = Integers{}
	ints := Integers{[]uint64{2, 1, 3, 4}}
	fmt.Println(sort.IsSorted(ints))

	sort.Sort(ints)
	fmt.Println("sorted", ints)

	rev := sort.Reverse(ints)
	sort.Slice(ints.data, func(i int, j int) bool {
		return ints.data[i]-ints.data[j] <= 0
	})
	fmt.Println("rev1", ints)

	rev = sort.Reverse(rev)
	sort.Slice(ints.data, rev.Less)

	fmt.Println("rev2", ints)

	fmt.Println(ints)
}
