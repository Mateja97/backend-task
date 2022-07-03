package rates

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type CoinGecko struct {
	Id         string     `json:"id,omitempty"`
	Symbol     string     `json:"symbol,omitempty"`
	MarketData MarketData `json:"market_data,omitempty"`
}
type HistoryResponse struct {
	ID     string
	Values HistoryValues
}
type HistoryValues struct {
	OnChain string
	OnGecko string
}
type MarketData struct {
	CurrentPrice Price `json:"current_price,omitempty"`
}
type Price map[string]float64

func GetCoin(id string) CoinGecko {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/" + id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("[ERROR] GetRates request failed", err)
	}
	defer resp.Body.Close()
	var conv CoinGecko
	if err := json.NewDecoder(resp.Body).Decode(&conv); err != nil {
		log.Fatal("Decoding rates failed")
	}

	return conv
}

func CoinHistory(hisResp HistoryResponse) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		date := vars["date"]

		layout := "02-01-2006"
		_, err := time.Parse(layout, date)

		if err != nil {
			log.Println("[ERROR]", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Wrong date format"))
		}
		url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/history?date=%s", id, date)
		log.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Println("[ERROR] CoinHistory request failed", err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("[ERROR] CoinHistory read alls failed", err)
		}
		var c CoinGecko
		if err := json.Unmarshal(body, &c); err != nil { // Parse []byte to go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}
		if err != nil {
			log.Println("[ERROR] Coin decode: ", err)
		}
		if err != nil {
			log.Println("[ERROR] Coin decode: ", err)
		}
		w.Write(body)
	}
}
