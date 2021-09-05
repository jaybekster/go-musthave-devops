package store

import (
	"strconv"
	"time"
)

type Store struct {
	data []map[string]interface{}
}

func (s *Store) Save(data interface{}) interface{} {
	item := make(map[string]interface{})

	key := strconv.FormatInt(time.Now().Unix(), 10)

	item[key] = data

	s.data = append(s.data, item)

	return item[key]
}

func (s *Store) Get() []map[string]interface{} {
	return s.data
}

func NewStore() *Store {
	return &Store{
		data: make([]map[string]interface{}, 0),
	}
}
