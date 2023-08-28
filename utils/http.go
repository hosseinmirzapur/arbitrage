package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetRequest(usdtPriceURL string) (map[string]interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, usdtPriceURL, nil)
	if err != nil {
		log.Fatalf("client: could not create request: %s\n", err)
	}
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("client: error making http request: %s\n", err)
	}

	fmt.Println("client: got a response!")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var jsonMap map[string]interface{}
	json.Unmarshal(data, &jsonMap)
	return jsonMap, nil
}
