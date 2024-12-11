package server

import (
	"net/http"
	"signaling-server/internal/handlers"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()

	signalingHandler := handlers.NewSignalingHandler()
	mux.HandleFunc("/ws", signalingHandler.HandleWebSocket)

	fs := http.FileServer(http.Dir("test"))
    mux.Handle("/test/", http.StripPrefix("/test/", fs))

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
