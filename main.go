package main

import (
	"net/http"
	v1 "upload/bussiness/v1"
)

func main() {
	http.HandleFunc("/v1/upload", v1.Upload)
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
