package main

import (
	"log"
	"net/http"
)

func main() {
	setupAPI()

	log.Fatal(http.ListenAndServe("127.1:3001", nil))
}

func setupAPI() {
	manager := NewManager()

	http.Handle("/", http.FileServer(http.Dir("./dist")))
	http.HandleFunc("/ws", manager.serveWs)
}
