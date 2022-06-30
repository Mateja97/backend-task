package server

import (
	socket "backend-task/backend/websocket"
	"log"
	"net/http"
)

func (s *Server) RegisterClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := socket.NewHandler(w, r)
		if err != nil {
			log.Println("[ERROR] Socket failed", err)
		}
		socket.Clients[ws] = true
	}
}
