package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	listeningHost := "localhost"
	listeningPort := "3000"
	listeningAddress := listeningHost + ":" + listeningPort
	log.Println("Listening on " + listeningAddress + "...")

	err := http.ListenAndServe(listeningAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
