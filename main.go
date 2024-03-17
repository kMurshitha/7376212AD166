package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ms.Murshitha!")
	})
	http.HandleFunc("/regnumber", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212AD166")
	})
	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "AI&DS")
	})
	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Purple")
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/hello\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
