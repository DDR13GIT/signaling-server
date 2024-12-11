package models

type SDPMessage struct {
	Type string `json:"type"`
	SDP  string `json:"sdp"`
}
