package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func baseGet(req *http.Request) (map[string]interface{}, error) {

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 && res.StatusCode < 400 {
		redirectUrl, err := res.Location()
		if err != nil {
			log.Fatal("Error getting redirect location:", err)
		}

		fmt.Println("client: redirecting to:", redirectUrl)

		req.URL = redirectUrl
		res, err = client.Do(req)
		if err != nil {
			log.Fatal("error sending redirect request: ", err)
		}
		defer res.Body.Close()
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
func GetRequest(url string) (map[string]interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("client: could not create request: %s\n", err)
	}
	return baseGet(req)
}

func AbanGetRequest(url string) (map[string]interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("client: could not create request: %s\n", err)
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("ABANTETHER_API_KEY"))

	return baseGet(req)
}
