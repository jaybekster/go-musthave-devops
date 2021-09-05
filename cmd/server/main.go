package main

import (
	"github.com/jaybekster/go-musthave-devops/internal/server"
	"github.com/jaybekster/go-musthave-devops/internal/store"
)

func main() {
	store := store.NewStore()
	srv := server.NewServer(store)

	srv.Serve()
}
