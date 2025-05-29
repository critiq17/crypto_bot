package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PriceResponse struct {
	Sybol string `json:"symbol"`
	Price string `json:"price"`
}

func GetPrice(symbol string) (string, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return " ", fmt.Errorf("binnace error code: %d", resp.StatusCode)
	}

	var price PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&price); err != nil {
		return "", err
	}

	return price.Price, nil
}
