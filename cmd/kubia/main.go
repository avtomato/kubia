package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Kubia server starting...")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("Received request from " + r.RemoteAddr)
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln("Error get hostname", err)
	}
	fmt.Fprintf(w, "You've hit %s \n", hostname)
}
