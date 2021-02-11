package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/upload", handler func(ResponseWriter, *Request))
	http.ListenAndServe()
}
