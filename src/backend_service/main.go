package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

var client RabbitClient

func main() {

	client.Connect()

	go client.Subscribe()
	http.Handle("/foo", apiHandler{})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q/n", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello foo")
	client.Publish()
}
