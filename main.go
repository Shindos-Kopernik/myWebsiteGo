package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/home/shindos//webSite")))
	http.ListenAndServe(":5000", nil)
}
