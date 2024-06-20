package main

import "fmt"

type Storage interface {
	Store(key string, value any)
}

type LocalStorage struct {
	store func(key string, value any)
}

func (ls LocalStorage) Store(key string, value any) {
	ls.store(key, value)

}

func main() {
	store := Storage(nil)

	localStorage := LocalStorage{}
	localStorage.store = storing

	store = localStorage
	store.Store("11", 12)
	fmt.Println("store is", store)
}

func storing(key string, value any) {
	fmt.Println(key)
}
