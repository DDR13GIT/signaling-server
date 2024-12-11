package main

import (
	"log"
	"signaling-server/server"
)

func main() {
	srv := server.NewServer()
	log.Printf("Starting signaling server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
