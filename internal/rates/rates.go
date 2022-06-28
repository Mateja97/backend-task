package rates

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Coin struct {
	MarketData MarketData `json:"market_data,omitempty"`
}
type MarketData struct {
	CurrentPrice Price `json:"current_price,omitempty"`
}
type Price map[string]float64

func GetCoin(id string) Coin {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/" + id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("[ERROR] GetRates request failed", err)
	}
	defer resp.Body.Close()
	var conv Coin
	if err := json.NewDecoder(resp.Body).Decode(&conv); err != nil {
		log.Fatal("Decoding rates failed")
	}

	return conv
}

func CoinHistory() http.HandlerFunc {
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
		url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/history&date=%s", id, date)
		log.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Println("[ERROR] CoinHistory request failed", err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("[ERROR]", err)
		}
		w.Write(data)
	}
}
