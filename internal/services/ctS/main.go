package main

import (
	"backend-task/internal/services/ctS/chaintracker"
	"flag"
	"log"
	"os"
	"os/signal"
)

var port = flag.String("port", ":8080", "")
var contract = flag.String("contract", "", "contract address")
var network = flag.String("goerli", "wss://eth-goerli.alchemyapi.io/v2/zsS8m3LQVW9mivXB_QRJh9xk2soIQfi9", "network url")
var serverURL = flag.String("server", "http://localhost:8081", "server url")

func main() {
	flag.Parse()
	ct := chaintracker.ChainTracker{}

	if err := ct.Init(*port, *contract, *network, *serverURL); err != nil {
		log.Println("[ERROR] Chain Tracker init failed ", err.Error())
	}
	go ct.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done
	ct.Stop()
}
