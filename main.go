package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const PORT = "8000"

var index = template.Must(template.ParseFiles("index.html"))

type Message struct {
	Text      string
	User      string
	TimeStamp string
}

var Messages []*Message

func tempHandler(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, Messages)
}

func addMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := r.Form.Get("user")
	text := r.Form.Get("text")
	timeStamp := time.Now().Format(time.RFC822)

	Messages = append(Messages, &Message{
		User:      user,
		Text:      text,
		TimeStamp: timeStamp,
	})

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))

	mux.HandleFunc("/", tempHandler)
	mux.HandleFunc("/add-message", addMessage)
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Printf("Listening on %s...", PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), mux)

	if err != nil {
		log.Fatalln("Server Error:", err)
	}
}
