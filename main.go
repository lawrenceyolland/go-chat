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

var messages []*Message

func tempHandler(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, nil)
	
	for _, msg := range messages {
		w.Write([]byte(fmt.Sprintf("<span>User:%s</span><br/><span>Message:%s</span><br/><span>Message:%s</span><br/>", msg.User, msg.Text, msg.TimeStamp)))
	}
}

func addMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := r.Form.Get("user")
	text := r.Form.Get("text")
	timeStamp := time.Now().Format(time.RFC822)

	messages = append(messages, &Message{user, text, timeStamp})

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", tempHandler)
	mux.HandleFunc("/add-message", addMessage)

	fmt.Printf("Listening on %s...", PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), mux)

	if err != nil {
		log.Fatalln("Server Error:", err)
	}
}
