package server

import (
	"net/http"
)

type server struct {
	store store
	srv   *http.Server
}

func NewServer(store Store) *server {
	return &server{store: store, srv: &http.Server{}}
}
