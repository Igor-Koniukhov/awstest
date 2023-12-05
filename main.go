package main

import (
	"fmt"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	fmt.Fprintf(w, "IP Address: %s", clientIP)

}

func main() {

	http.HandleFunc("/", ApiHandler)

	http.ListenAndServe(":8080", nil)

}
