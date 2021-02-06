package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	if err := http.ListenAndServe(":4000", server); err != nil {
		log.Fatalf("could not listen on port 4000 %v", err)
	}
}
