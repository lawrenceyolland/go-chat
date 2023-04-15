package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

const PORT = "8000"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Text string `json:"text"`
	// TimeStamp string
}

var messages []websocket.Conn

func handleHomePath(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	messages = append(messages, *conn)

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.Fatal(err)
		}

		for _, message := range messages {
			m := &Message{Text: string(msg)}

			if err = message.WriteJSON(m); err != nil {
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/", handleHomePath)
	http.HandleFunc("/chat", handleMessages)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)

	if err != nil {
		log.Fatalln("Server Error:", err)
	}
}
