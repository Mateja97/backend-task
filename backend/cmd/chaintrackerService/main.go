package main

import (
	"backend-task/backend/chaintracker"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
)

var port = flag.String("port", ":8080", "")
var contract = flag.String("contract", "", "contract address")
var network = flag.String("goerli", "wss://eth-goerli.alchemyapi.io/v2/zsS8m3LQVW9mivXB_QRJh9xk2soIQfi9", "network url")

//var serverURL = flag.String("server", "http://localhost:8081", "server url")
var kafkaBrokers = flag.String("kafka.brokers", "", "kafka brokersl")
var kafkaTopic = flag.String("kafka.topic", "chain", "kafka topic")

func main() {
	flag.Parse()
	ct := chaintracker.ChainTracker{}
	brokers := strings.Split(*kafkaBrokers, ",")
	if err := ct.Init(*port, *contract, *network, *kafkaTopic, brokers); err != nil {
		log.Println("[ERROR] Chain Tracker init failed ", err.Error())
	}
	go ct.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done
	ct.Stop()
}
