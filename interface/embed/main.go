package main

type Hasher interface {
	Hash() []byte
	Write(data []byte)
}

// can't have an implementation on interface
func (hash *Hasher) Hash() []byte {
	return []byte("hash")
}

type FixedHasher interface {
	Hasher
	// Hash() int; // error, dupicate stuff without different params as embedded interface
	Hash() []byte // no error, dupicate stuff with same params as embedded interface
}

type Sha struct {
	FixedHasher
}

func main() {
	var sha Hasher = Sha{}

	sha.Write([]byte{1, 2, 3})
}
