package store

import (
	"strconv"
	"time"
)

type Store struct {
	data []map[string]interface{}
}

func (s *Store) save(data interface{}) interface{} {
	item := make(map[string]interface{})

	key := strconv.FormatInt(time.Now().Unix(), 10)

	item[key] = data

	s.data = append(s.data, item)

	return item[key]
}

func (s *Store) get() []map[string]interface{} {
	return s.data
}
