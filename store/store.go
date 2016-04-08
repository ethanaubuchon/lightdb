package store

import(
   "errors"
)

type Store struct {
   data map[string]string
}

func (store *Store) Get(key string) (value string, err error) {
   value, ok := store.data[key]

   if ok {
      return value, nil
   }
   return "", errors.New("No value set")
}

func (store *Store) Set(key string, value string) {
   if store.data == nil {
      store.Init()
   }
   store.data[key] = value
}

func (store *Store) Unset(key string) {
   delete(store.data, key)
}

func (store *Store) Init() {
   store.data = make(map[string]string)
}
