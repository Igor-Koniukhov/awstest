package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handler worked - send IP.")

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		if ip != nil && ip.IsGlobalUnicast() {
			fmt.Fprintf(w, "IP Address: %s", ip)
		}
	}

}

func main() {
	log.Println("The API has started.")

	http.HandleFunc("/", ApiHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
