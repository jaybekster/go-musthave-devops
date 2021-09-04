package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		m, _ := url.ParseQuery(r.URL.RawQuery)

		metricId := m["id"]
		metricType := m["type"]
		metricValue := m["value"]

		fmt.Println(metricId, metricType, metricValue)
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", mainHandler)

	log.Fatal(http.ListenAndServe(":8880", nil))
}
