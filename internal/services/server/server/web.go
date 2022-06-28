package server

import (
	"backend-task/internal/chain"
	"backend-task/internal/socket"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *Server) ChainValue() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var entity chain.ChainEntity
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&entity)
		if err != nil {
			panic(err)
		}
		text := fmt.Sprintf("%s: $%s", entity.Symbol, entity.Amount.String())
		log.Println(text)
		s.chainCh <- entity
	}
}
func (s *Server) ChainPriceChange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := socket.NewHandler(w, r)
		if err != nil {
			log.Println("[ERROR] Socket failed", err)
		}
		socket.Clients[ws] = true
	}
}
