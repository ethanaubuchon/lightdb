package store

import (
	"errors"
	"sync"
)

type Storer interface {
	Get(key string) (value string, err error)
	Set(key string, value string)
	Unset(key string)
}

type Store struct {
	sync.RWMutex
	data map[string]string
}

func NewStore() *Store {
	return &Store{data: make(map[string]string)}
}

func (store *Store) Get(key string) (value string, err error) {
	store.RLock()
	value, ok := store.data[key]
	store.RUnlock()

	if ok {
		return value, nil
	}
	return "", errors.New("No value set")
}

func (store *Store) Set(key string, value string) {
	store.Lock()
	store.data[key] = value
	store.Unlock()
}

func (store *Store) Unset(key string) {
	store.Lock()
	delete(store.data, key)
	store.Unlock()
}
