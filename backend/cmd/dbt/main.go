package main

import (
	"backend-task/backend/cmd/dbt/server"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
)

var port = flag.String("port", ":8081", "server port")
var wsport = flag.String("port.ws", ":8082", "websocket port")
var publisher = flag.String("publisher.url", "http://localhost:8083", "publisher url")
var kafkaBrokers = flag.String("kafka.brokers", "", "kafka brokersl")
var kafkaTopic = flag.String("kafka.topic", "chain", "kafka topic")
var tokenIds = flag.String("token.ids", "ethereum,bitcoin,cardano", "values of the tokens that want to track")

func main() {
	flag.Parse()
	s := server.Server{}
	brokers := strings.Split(*kafkaBrokers, ",")
	ids := strings.Split(*tokenIds, ",")
	err := s.Init(*port, *wsport, *publisher, *kafkaTopic, brokers, ids)
	if err != nil {
		log.Println("[ERROR] Server init failed")
	}
	go s.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done
	s.Stop()
}
