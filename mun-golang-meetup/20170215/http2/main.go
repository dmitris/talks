package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":10443", "localhost.cert", "localhost.key", nil) // HL
	log.Fatal(err)
}
