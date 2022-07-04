package rates

type CoinGecko struct {
	Id         string     `json:"id,omitempty"`
	Symbol     string     `json:"symbol,omitempty"`
	MarketData MarketData `json:"market_data,omitempty"`
}
type CoinResponse struct {
	ID     string     `json:"id,omitempty"`
	Values CoinValues `json:"values,omitempty"`
}
type CoinValues struct {
	OnChain string `json:"on_chain,omitempty"`
	OnGecko string `json:"on_gecko,omitempty"`
}
type MarketData struct {
	CurrentPrice Price `json:"current_price,omitempty"`
}
type Price map[string]float64
