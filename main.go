package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "https://api.currencyapi.com/v3/latest?&base_currency=SEK&currencies=USD,AED,NOK"
	method := "GET"

	// Read the API key from the file
	apiKey, err := os.ReadFile("api-key.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("apikey", string(apiKey))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
