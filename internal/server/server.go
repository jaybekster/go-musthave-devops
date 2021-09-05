package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaybekster/go-musthave-devops/internal/store"
)

type server struct {
	store *store.Store
	srv   *http.Server
}

func NewServer(store *store.Store) *server {
	return &server{store: store, srv: &http.Server{}}
}

func (s *server) Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", saveHandler(s)).Methods("POST")
	r.HandleFunc("/", getHandler(s)).Methods("GET")

	s.srv = &http.Server{
		Handler: r,
		Addr:    ":" + "8080",
	}

	log.Printf("Server start lisening on ip")

	log.Fatal(s.srv.ListenAndServe())
}
