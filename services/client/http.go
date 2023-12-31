package client

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetRequest(url string) (map[string]interface{}, error) {
	agent := fiber.Get(url)
	_, body, errs := agent.Bytes()

	if len(errs) > 0 {
		return nil, errors.New("client: could not make a request to the exchange")
	}

	var data map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func PostRequest(url string, body []byte) (map[string]interface{}, error) {
	agent := fiber.Post(url)

	agent.Body(body)
	agent.Request().Header.Add("Accept", "application/json")
	agent.Request().Header.Add("Content-Type", "application/json")

	_, receivedData, errs := agent.Bytes()
	if len(errs) > 0 {
		return nil, errors.New("client: could not make a request to the exchange")
	}

	var data map[string]interface{}
	err := json.Unmarshal(receivedData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
