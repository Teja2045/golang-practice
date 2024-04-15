package main

import (
	"errors"
	"fmt"
)

var ErrNotFound error = errors.New("Not found")

func main() {
	err := throwNewError()
	if err != nil {
		if err == ErrNotFound {
			fmt.Println("defined error", err)
		} else {
			fmt.Println("unknown error", err)
		}
	}

	err = throwDefinedError()
	if err != nil {
		if err == ErrNotFound {
			fmt.Println("defined error", err)
		} else {
			fmt.Println("unknown error", err)
		}
	}
}

func throwNewError() error {
	return errors.New("Not found")
}

func throwDefinedError() error {
	return ErrNotFound
}
