package chain

import (
	"encoding/json"
	"log"
	"math/big"

	"github.com/Shopify/sarama"
	"github.com/gorilla/websocket"
)

type ChainCache struct {
	//coin/amount
	ChainValues map[string]string
	//ex: coin/date/amount
	ChainValuesHistory map[string]map[string]string
}

func (cc *ChainCache) Init() {
	cc.ChainValues = make(map[string]string)
	cc.ChainValuesHistory = make(map[string]map[string]string)
}

type ChainEntity struct {
	ID     string   `json:"id,omitempty"`
	Symbol string   `json:"symbol,omitempty"`
	Amount *big.Int `json:"amount,omitempty"`
	Date   string
}

//WriteChainCacheToClients Sends ChainValues to the websocket for specified clients
func (cc *ChainCache) WriteChainCacheToClients(clients map[*websocket.Conn]bool) {
	for ws := range clients {
		err := ws.WriteJSON(cc.ChainValues)
		if err != nil {
			log.Println("[ERROR] Sending json: ", err)
			ws.Close()
		}
	}
}
func (cc *ChainCache) StoreCache(e *sarama.ConsumerMessage) {
	var entity ChainEntity
	err := json.Unmarshal(e.Value, &entity)
	if err != nil {
		log.Println("[ERROR] Unmarshaling entity failed", err)
		return
	}
	if entity == (ChainEntity{}) {
		return
	}
	if _, ok := cc.ChainValuesHistory[entity.Symbol]; !ok {
		cc.ChainValuesHistory[entity.Symbol] = make(map[string]string)
	}
	cc.ChainValues[entity.Symbol] = entity.Amount.String()
	cc.ChainValuesHistory[entity.Symbol][entity.Date] = entity.Amount.String()
	log.Println("[INFO] Stored from kafka to the cache", entity.Symbol, " ", cc.ChainValues[entity.Symbol])

}
