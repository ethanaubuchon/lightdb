package store

import (
	"errors"
	"sync"
)

type Storer interface {
	Get(key string) (value interface{}, err error)
	Set(key string, value interface{})
	Unset(key string)
	Begin() *Tx
}

type Store struct {
	sync.RWMutex
	data map[string]interface{}
}

func NewStore() *Store {
	return &Store{data: make(map[string]interface{})}
}

func (store *Store) Get(key string) (value interface{}, err error) {
	store.RLock()
	value, ok := store.data[key]
	store.RUnlock()

	if ok {
		return value, nil
	}
	return "", errors.New("No value set")
}

func (store *Store) Set(key string, value interface{}) {
	store.Lock()
	store.data[key] = value
	store.Unlock()
}

func (store *Store) Unset(key string) {
	store.Lock()
	delete(store.data, key)
	store.Unlock()
}

func (store *Store) Begin() *Tx {
	return &Tx{
		parent:    store,
		overwrite: make(map[string]interface{}),
	}
}

// Transaction

type Tx struct {
	parent    Storer
	overwrite map[string]interface{}
}

func (tx *Tx) Get(key string) (value interface{}, err error) {
	data, ok := tx.overwrite[key]
	if ok {
		return data, nil
	}

	return tx.parent.Get(key)
}

func (tx *Tx) Set(key string, value interface{}) {
	tx.overwrite[key] = value
}

func (tx *Tx) Unset(key string) {
	tx.overwrite[key] = nil
}

func (tx *Tx) Begin() *Tx {
	return &Tx{
		parent:    tx,
		overwrite: make(map[string]interface{}),
	}
}

func (tx *Tx) Rollback() {
	tx = nil
}

func (tx *Tx) Commit() {
	for key, value := range tx.overwrite {
		tx.parent.Set(key, value)
	}
	tx = nil
}
