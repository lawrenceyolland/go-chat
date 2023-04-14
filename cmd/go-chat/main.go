package main

import (
	// "log"
	"fmt"
	"log"
	"net/http"
)

const PORT = "8000"

func tempHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	message := fmt.Sprintf("<h1>Hi! Welcome to %s</h1>", path)
	w.Write([]byte(message))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", tempHandler)

	fmt.Printf("Listening on %s...", PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), mux)

	if err != nil {
		log.Fatalln("Server Error:", err)
	}
}
