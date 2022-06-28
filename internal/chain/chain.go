package chain

import (
	"backend-task/internal/socket"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"net/http"

	"github.com/gorilla/websocket"
)

type ChainEntity struct {
	ID     string   `json:"id,omitempty"`
	Symbol string   `json:"symbol,omitempty"`
	Amount *big.Int `json:"amount,omitempty"`
}

func ChainPriceChange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := socket.NewHandler(w, r)
		if err != nil {
			log.Println("[ERROR] Socket failed", err)
		}
		var entity ChainEntity
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&entity)
		if err != nil {
			panic(err)
		}
		text := fmt.Sprintf("%s: $%s", entity.Symbol, entity.Amount.String())
		ws.WriteMessage(websocket.TextMessage, []byte(text))
	}
}
