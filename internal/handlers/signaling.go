package handlers

import (
	"log"
	"net/http"
	"signaling-server/internal/models"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // In production, implement proper origin checking
	},
}

type SignalingHandler struct {
	mu    sync.Mutex
	peers []*websocket.Conn
}

func NewSignalingHandler() *SignalingHandler {
	return &SignalingHandler{
		peers: make([]*websocket.Conn, 0, 2),
	}
}

func (h *SignalingHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	log.Printf("New peer connected. Total peers: %d", len(h.peers)+1)
	h.mu.Lock()
	if len(h.peers) >= 2 {
		h.mu.Unlock()
		conn.Close()
		return
	}
	h.peers = append(h.peers, conn)
	h.mu.Unlock()

	defer func() {
		h.mu.Lock()
		for i, peer := range h.peers {
			if peer == conn {
				h.peers = append(h.peers[:i], h.peers[i+1:]...)
				break
			}
		}
		h.mu.Unlock()
		conn.Close()
	}()

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		log.Printf("Received message type: %s", msg.Type)
		// Forward message to the other peer
		h.mu.Lock()
		for _, peer := range h.peers {
			if peer != conn {
				err := peer.WriteJSON(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
			}
		}
		h.mu.Unlock()
	}
}
