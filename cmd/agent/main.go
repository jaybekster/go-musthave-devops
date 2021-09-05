package main

import "net/http"

func main() {
	client := http.Client{}

	response, err := client.Post("//localhost:8880")
}
