package server

import (
	"backend-task/backend/rates"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

//CoinGeckoTracker track values from coin gecko for specified token ids and update chain values
func (s *Server) CoinGeckoTracker() {
	for {
		<-s.ticker.C
		//Track all specified token ids
		for _, id := range s.tokenIds {
			url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", id)
			resp, err := http.Get(url)
			if err != nil {
				log.Println("[ERROR] CoinValues request failed", err)
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println("[ERROR] CoinValues read alls failed", err)
			}
			var c rates.CoinGecko
			if err := json.Unmarshal(body, &c); err != nil { // Parse []byte to go struct pointer
				fmt.Println("[ERROR]Unmarshal CoinGecko failed", err)
			}
			price := c.MarketData.CurrentPrice["usd"]
			if _, ok := s.cc.ChainValues[c.Symbol]; !ok {
				log.Println("[INFO] Value for ", c.Symbol, " not found")
				s.cc.ChainValues[c.Symbol] = ""
				continue
			}
			priceString := fmt.Sprintf("%d", int(price*100))
			val, err := strconv.Atoi(priceString)
			if err != nil {
				log.Println("[ERROR] Convert val failed", err)
				continue
			}
			req := struct {
				Symbol string   `json:"symbol,omitempty"`
				Amount *big.Int `json:"amount,omitempty"`
			}{
				Symbol: c.Symbol,
				Amount: big.NewInt(int64(val)),
			}
			req_json, err := json.Marshal(req)

			if err != nil {
				log.Fatal(err)
			}
			//Check if coing gecko price is 2% differs from the on chain value
			if price*100 > float64(val*2/100) {
				resp, err := http.Post(s.publisherUrl+"/publish", "application/json", bytes.NewBuffer(req_json))
				if err != nil || resp.StatusCode != 200 {
					log.Println("[ERROR] Sending data to publisher failed: ", err)
					continue
				}
			}
			log.Println("[INFO] Data sent to the publisher for symbol: ", c.Symbol, " price:", priceString)
		}
	}
}
