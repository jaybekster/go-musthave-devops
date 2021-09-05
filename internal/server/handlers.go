package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type QueryRequest struct {
	ID    string
	Type  string
	Value string
}

func saveHandler(s *server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)

		metricId := r.FormValue("id")
		metricType := r.FormValue("type")
		metricValue := r.FormValue("value")

		qr := QueryRequest{
			ID:    metricId,
			Type:  metricType,
			Value: metricValue,
		}

		log.Printf("%+v", qr)

		s.store.Save(qr)

		fmt.Println(metricId, metricType, metricValue)

		w.WriteHeader(http.StatusOK)
	}
}

func getHandler(s *server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := s.store.Get()

		jsonResponse, jsonError := json.Marshal(data)

		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}
