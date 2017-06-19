package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/contact", serveContact)
	http.ListenAndServe(":8888", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func serveContact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello to contact page"))
}
