package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World.")
		fmt.Fprintln(w, "My message2.")
	})
	log.Fatalf("error: %s", http.ListenAndServe(":8080", nil))
	
}
