package server

import (
	"backend-task/internal/chain"
	"backend-task/internal/rates"
	"backend-task/internal/socket"
	"backend-task/internal/util"
	"fmt"

	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	webSocket   *http.Server
	s           *http.Server
	ticker      *time.Ticker
	chainCh     chan chain.ChainEntity
	ChainValues map[string]string
}

func (s *Server) Init(port, wsPort string) error {
	s.ticker = time.NewTicker(60 * time.Second)
	s.ChainValues = make(map[string]string)
	s.chainCh = make(chan chain.ChainEntity)
	r := mux.NewRouter()
	r.HandleFunc("/price/{id}/history/{date}", rates.CoinHistory()).Methods("GET")
	r.HandleFunc("/price/chain", s.ChainValue()).Methods("POST")

	rWS := mux.NewRouter()
	rWS.HandleFunc("/price/change", s.ChainPriceChange())

	s.s = &http.Server{
		Addr:    port,
		Handler: util.CORSEnabled(r),
	}
	s.webSocket = &http.Server{
		Addr:    wsPort,
		Handler: util.CORSEnabled(rWS),
	}
	return nil
}

func (s *Server) Run() {
	go func() {
		if err := s.s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("[ERROR] Server shutdown with error:", err.Error())
		}
	}()
	go func() {
		if err := s.webSocket.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("[ERROR] WebSocket shutdown with error:", err.Error())
		}
	}()

	go func() {
		for {
			e := <-s.chainCh
			s.ChainValues[e.Symbol] = e.Amount.String()

			log.Println(s.ChainValues)
			for ws := range socket.Clients {
				err := ws.WriteJSON(s.ChainValues)

				if err != nil {
					log.Println("[ERROR] Sending json: ", err)
					ws.Close()
				}
			}
		}

	}()

	for {
		<-s.ticker.C
		btc := rates.GetCoin("bitcoin")
		eth := rates.GetCoin("ethereum")
		btcPrice := btc.MarketData.CurrentPrice["usd"]
		ethPrice := eth.MarketData.CurrentPrice["usd"]

		fmt.Printf("btc: $%.2f$", btcPrice)
		fmt.Println()
		fmt.Printf("eth: $%.2f", ethPrice)
		fmt.Println()
	}

}
func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.s.Shutdown(ctx)
	if err != nil {
		log.Fatalln("[ERROR] Could not stop server gracefully:", err.Error())
	}
	err = s.webSocket.Shutdown(ctx)
	if err != nil {
		log.Fatalln("[ERROR] Could not stop websocket gracefully:", err.Error())
	}
	s.ticker.Stop()
	log.Println("Graceful shutdown complete.")
}
