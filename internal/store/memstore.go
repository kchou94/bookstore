package store

import (
	mystore "bookstore/store"
	factory "bookstore/store/factory"
	"sync"
)

type MemStore struct {
	sync.Mutex
	books map[string]*mystore.Book
}

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}
