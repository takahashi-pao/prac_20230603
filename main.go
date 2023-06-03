package main

import (
	"fmt"
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		name := queryParams.Get("name")
		handlerHello(w, r, name)
	})
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintln(w, count)
}

func handlerHello(w http.ResponseWriter, r *http.Request, name string) {
	/*
		message := strings.Replace("Hello, {name}", "{name}", name, 1)
	*/

	fmt.Fprintf(w, "Hello, %s\n", name)
}
