package server

import (
	"backend-task/backend/rates"
	socket "backend-task/backend/websocket"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//RegisterClient registers client to the web socket
func (s *Server) RegisterClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := socket.NewHandler(w, r)
		if err != nil {
			log.Println("[ERROR] Socket failed", err)
		}
		socket.Clients[ws] = true
	}
}

//CoinHistory returns values onchain and on coingecko for specified coin at the specified date
func (s *Server) CoinHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		date := vars["date"]
		layout := "02-01-2006"

		reqTime, err := time.Parse(layout, date)
		if err != nil {
			log.Println("[ERROR]", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Wrong date format"))
			return
		}
		url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/history?date=%s", id, date)
		resp, err := http.Get(url)
		if err != nil {
			log.Println("[ERROR] CoinHistory request failed", err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("[ERROR] CoinHistory read alls failed", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var c rates.CoinGecko
		if err := json.Unmarshal(body, &c); err != nil { // Parse []byte to go struct pointer
			fmt.Println("[ERROR]Unmarshal CoinGecko failed", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t := time.Date(reqTime.Year(), reqTime.Month(), reqTime.Day(), 0, 0, 0, 0, time.UTC)

		coinGecko := fmt.Sprintf("%.2f", c.MarketData.CurrentPrice["usd"])
		chainValue := s.cc.ChainValuesHistory[c.Symbol][t.Format(layout)]
		val, err := strconv.Atoi(chainValue)
		if err != nil {
			log.Println("[ERROR] Convert val failed", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		stringVal := fmt.Sprintf("%.2f", float64(val/100))
		HistoryResponse := rates.CoinResponse{
			ID: id,
			Values: rates.CoinValues{
				OnChain: stringVal,
				OnGecko: coinGecko,
			},
		}
		data, err := json.Marshal(HistoryResponse)
		if err != nil {
			log.Println("[ERROR] Marshaling history responsefailed: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		w.Write(data)
	}
}

//CoinValues returns coin values from chain and coingecko for the specified id
func (s *Server) CoinValues() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", id)
		resp, err := http.Get(url)
		if err != nil {
			log.Println("[ERROR] Coingecko request failed", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("[ERROR] Cannot read coingecko response", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var c rates.CoinGecko
		if err := json.Unmarshal(body, &c); err != nil { // Parse []byte to go struct pointer
			fmt.Println("[ERROR]Unmarshal CoinGecko failed", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		coinGecko := fmt.Sprintf("%.2f", c.MarketData.CurrentPrice["usd"])
		chainValue := s.cc.ChainValues[c.Symbol]

		val, err := strconv.Atoi(chainValue)
		if err != nil {
			log.Println("[ERROR] Convert val failed", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//On chain values are *100
		stringVal := fmt.Sprintf("%.2f", float64(val/100))

		valuesResponse := rates.CoinResponse{
			ID: id,
			Values: rates.CoinValues{
				OnChain: stringVal,
				OnGecko: coinGecko,
			},
		}
		data, err := json.Marshal(valuesResponse)
		if err != nil {
			log.Println("[ERROR] Marshaling history responsefailed: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		w.Write(data)

	}
}
