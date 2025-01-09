# Signaling Server

This Go-based signaling server facilitates WebRTC peer-to-peer connections by managing the exchange of Session Description Protocol (SDP) messages and Interactive Connectivity Establishment (ICE) candidates between two clients.

## Features

- **WebSocket Connectivity**: Clients can establish a WebSocket connection to the server for real-time signaling.
- **Peer-to-Peer Signaling**: Enables two clients to exchange SDP and ICE candidates, essential for establishing a direct WebRTC connection.

## Prerequisites

- **Go**: Ensure that Go is installed on your system. You can download it from the [official website](https://golang.org/dl/).

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/DDR13GIT/signaling-server.git
   cd signaling-server
   ```
2. **Install Dependencies**:

   ```bash
   go mod tidy
   ```
3. **Start the server**:

   ```bash
   go run cmd/main.go
   ```
By default, the server listens on port 8080.

4. **Connect Clients**:

Clients can connect to the server's WebSocket endpoint at ws://localhost:8080/ws.
Once connected, clients can exchange SDP offers/answers and ICE candidates through the server to establish a direct peer-to-peer connection.    
