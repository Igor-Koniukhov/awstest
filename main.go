package main

import (
	"fmt"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handler worked - send IP.")
	ip := os.Getenv("EC2")
	fmt.Fprintf(w, "IP Address: %s", ip)

}

func main() {
	log.Println("The API has started.")

	http.HandleFunc("/", ApiHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
