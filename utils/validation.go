package utils

import (
	"encoding/json"
	"net/http"
)

func IsValidBreed(breed string) bool {
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var breeds []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&breeds)

	for _, b := range breeds {
		if b["name"] == breed {
			return true
		}
	}
	return false
}
