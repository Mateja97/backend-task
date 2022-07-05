package chain

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//ChainValue rest api to receive changes from the chain and send received entity to the specified channel
func (cc *ChainCache) ChainValue(ch chan ChainEntity) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var entity ChainEntity
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&entity)
		if err != nil {
			log.Println("[ERROR] ChainValue Decoding failed", err)
			return
		}
		text := fmt.Sprintf("%s: $%s", entity.Symbol, entity.Amount.String())
		log.Println(text)
		cc.Lock()
		cc.ChainValues[entity.Symbol] = entity.Amount.String()
		cc.Unlock()
		ch <- entity
	}
}
