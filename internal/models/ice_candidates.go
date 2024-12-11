package models

type ICECandidate struct {
    Candidate     string  `json:"candidate"`
    SDPMLineIndex uint16  `json:"sdpMLineIndex"`
    SDPMid        string  `json:"sdpMid"`
}