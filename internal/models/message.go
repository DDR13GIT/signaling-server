package models

type MessageType string

const (
    Offer        MessageType = "offer"
    Answer       MessageType = "answer"
    IceCandidate MessageType = "ice-candidate"
)

type Message struct {
    Type      MessageType     `json:"type"`
    SDP       *SDPMessage     `json:"sdp,omitempty"`
    Candidate *ICECandidate   `json:"candidate,omitempty"`
}