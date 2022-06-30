package main

import (
	"backend-task/backend/chainpublisher"
	"flag"
	"log"
	"os"
	"os/signal"
)

var port = flag.String("port", ":8083", "server port")
var network = flag.String("goerli", "wss://eth-goerli.alchemyapi.io/v2/zsS8m3LQVW9mivXB_QRJh9xk2soIQfi9", "network url")
var contract = flag.String("contract", "", "contract address")
var privateKey = flag.String("private.key", "", "private key")

func main() {
	flag.Parse()
	cp := chainpublisher.ChainPublisher{}

	err := cp.Init(*network, *privateKey, *contract, *port)
	if err != nil {
		log.Fatalln("[ERROR] Chainpublisher init failed")
	}
	go cp.Run()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done
	cp.Stop()
}
