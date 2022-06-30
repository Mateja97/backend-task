package server

import (
	"backend-task/backend/chain"
	"backend-task/backend/cors"
	"backend-task/backend/kafka"
	"backend-task/backend/rates"
	socket "backend-task/backend/websocket"
	"encoding/json"
	"fmt"

	"context"
	"log"
	"net/http"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
)

type Server struct {
	webSocket     socket.WebSocket
	s             *http.Server
	consumer      kafka.KafkaConsumer
	ticker        *time.Ticker
	cc            chain.ChainCache
	chEntity      chan chain.ChainEntity
	kafkaMessages chan *sarama.ConsumerMessage
	publisherUrl  string
}

func (s *Server) Init(port, wsPort, publisher, topic string, brokers []string) error {
	s.ticker = time.NewTicker(60 * time.Second)
	s.chEntity = make(chan chain.ChainEntity)
	s.kafkaMessages = make(chan *sarama.ConsumerMessage)
	s.publisherUrl = publisher
	//Kafka consumer init
	err := s.consumer.Init(brokers, topic)
	if err != nil {
		return err
	}
	s.cc.Init() //ChainCache init

	//Server for rest api
	r := mux.NewRouter()
	r.HandleFunc("/price/{icmd/chaintrackerd}/history/{date}", rates.CoinHistory()).Methods("GET")
	r.HandleFunc("/price/chain", s.cc.ChainValue(s.chEntity)).Methods("POST")
	s.s = &http.Server{
		Addr:    port,
		Handler: cors.CORSEnabled(r),
	}
	//Web socket to publish chain changes to the frontend
	rWS := mux.NewRouter()
	rWS.HandleFunc("/price/change", s.RegisterClient())

	s.webSocket.Init(wsPort, cors.CORSEnabled(rWS))
	return nil
}

func (s *Server) Run() {
	go func() {
		if err := s.s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("[ERROR] Server shutdown with error:", err.Error())
		}
	}()
	go func() {
		if err := s.webSocket.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("[ERROR] WebSocket shutdown with error:", err.Error())
		}
	}()

	go s.consumer.ConsumeMessage(s.kafkaMessages)

	go func() {
		for {
			e := <-s.kafkaMessages //Waiting for onchain change
			s.SentWS(e)
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
	err = s.webSocket.Stop(ctx)
	if err != nil {
		log.Fatalln("[ERROR] Could not stop websocket gracefully:", err.Error())
	}
	err = s.consumer.Stop()
	if err != nil {
		log.Fatalln("[ERROR] Could not stop consumer gracefully:", err.Error())
	}
	s.ticker.Stop()
	log.Println("Graceful shutdown complete.")
}

func (s *Server) SentWS(e *sarama.ConsumerMessage) {
	var entity chain.ChainEntity
	err := json.Unmarshal(e.Value, &entity)
	if err != nil {
		log.Println("[ERROR] Unmarshaling entity failed", err)
		return
	}
	if entity == (chain.ChainEntity{}) {
		return
	}
	s.cc.ChainValues[entity.Symbol] = entity.Amount.String()
	s.cc.WriteChainCacheToClients(socket.Clients)
}
