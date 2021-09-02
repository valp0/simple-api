package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	type Content struct {
		Message string `json:"message"`
	}

	type Body struct {
		Content `json:"data"`
	}

	w.Header().Set("Content-Type", "application/json")

	// We create our response body.
	resp := Body{Content{"Please use either /hello or /btc2usd endpoints."}}

	// Parsing response struct to JSON.
	jRes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, string(jRes))
	if err != nil {
		fmt.Println(err)
		return
	}
}
