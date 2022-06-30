package chain

import (
	"log"
	"math/big"
	"sync"

	"github.com/gorilla/websocket"
)

type ChainCache struct {
	ChainValues map[string]string
	*sync.RWMutex
}

func (cc *ChainCache) Init() {
	cc.ChainValues = make(map[string]string)
}

type ChainEntity struct {
	ID     string   `json:"id,omitempty"`
	Symbol string   `json:"symbol,omitempty"`
	Amount *big.Int `json:"amount,omitempty"`
}

func (cc *ChainCache) WriteChainCacheToClients(clients map[*websocket.Conn]bool) {
	for ws := range clients {
		err := ws.WriteJSON(cc.ChainValues)
		if err != nil {
			log.Println("[ERROR] Sending json: ", err)
			ws.Close()
		}
	}
}
