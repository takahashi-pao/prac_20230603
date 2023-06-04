package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		/*
			command => curl 'localhost:8080/greet?name={name}}'
		*/

		queryParams := r.URL.Query()
		name := queryParams.Get("name")
		handlerHello(w, r, name)
	})

	http.HandleFunc("/file", requestHtmlFileOnServer)
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

func requestHtmlFileOnServer(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://takahashi-pao.github.io/oretachi-omaetachi/index.html")
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
	}

	fmt.Fprintf(w, string(body))
}
