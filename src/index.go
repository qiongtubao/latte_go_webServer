package main

import (
	"net/http"
	"time"
	"server"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	server := server.CreateServer()
	server.Get("/time", timeHandler)
	server.Start()
}