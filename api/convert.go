package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetConvert(from string, to string, amout float64) (float64, error) {

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", from, to)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var result map[string]map[string]float64
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}
	price := result[from][to]
	return price * amout, nil
}
