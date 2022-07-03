package chain

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

		cc.ChainValues[entity.Symbol] = entity.Amount.String()
		ch <- entity
	}
}
