package main

import (
	"backend-task/internal/services/server/server"
	"flag"
	"log"
	"os"
	"os/signal"
)

var port = flag.String("port", ":8081", "server port")
var wsport = flag.String("port.ws", ":8082", "websocket port")

func main() {
	flag.Parse()
	s := server.Server{}
	err := s.Init(*port, *wsport)
	if err != nil {
		log.Println("[ERROR] Server init failed")
	}
	go s.Run()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	<-done
	s.Stop()
}
