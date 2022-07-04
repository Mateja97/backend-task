package server

import (
	"backend-task/backend/chain"
	"backend-task/backend/cors"
	"backend-task/backend/kafka"
	socket "backend-task/backend/websocket"

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
	tokenIds      []string
}

func (s *Server) Init(port, wsPort, publisher, topic string, brokers, ids []string) error {
	s.ticker = time.NewTicker(60 * time.Second)
	s.chEntity = make(chan chain.ChainEntity)
	s.kafkaMessages = make(chan *sarama.ConsumerMessage)
	s.publisherUrl = publisher
	s.tokenIds = ids
	//Kafka consumer init
	err := s.consumer.Init(brokers, topic)
	if err != nil {
		return err
	}
	s.cc.Init() //ChainCache init

	//Server for rest api
	r := mux.NewRouter()
	r.HandleFunc("/price/{id}/history/{date}", s.CoinHistory()).Methods("GET")
	r.HandleFunc("/price/{id}", s.CoinValues()).Methods("GET")
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
			//Waiting for onchain change to be published to the kafka
			e := <-s.kafkaMessages
			//Store chain values from kafka message to the cache
			s.cc.StoreCache(e)
			//Send te chain values to websocket client
			s.cc.WriteChainCacheToClients(socket.Clients)

		}
	}()
	s.CoinGeckoTracker()
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
