package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Bitcoin(w http.ResponseWriter, r *http.Request) {
	// Coinbase BTC to USD exchange rate endpoint.
	url := "https://api.coinbase.com/v2/prices/BTC-USD/buy?quote=true"

	// Calling GET HTTP method, which returns a response and an error.
	resp, err := http.Get(url)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Printf("%s is DOWN!\n", url)
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		// Closing response body if there is one.
		defer resp.Body.Close()

		// Reading response if HTTP request was successful.
		if resp.StatusCode == 200 {
			w.WriteHeader(http.StatusOK)
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			_, err := fmt.Fprintf(w, string(bodyBytes))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("%s -> Status Code: %d  \n", url, resp.StatusCode)
		}
	}
}
