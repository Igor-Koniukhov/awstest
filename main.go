package main

import (
	"fmt"
	"log"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	fmt.Fprintf(w, "IP Address: %s", clientIP)
	log.Println("Handler worked - send IP.")

}

func main() {
	log.Println("The API has started.")

	http.HandleFunc("/", ApiHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
