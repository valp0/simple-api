package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/valp0/simple-api/utils"
)

func TestBTC(t *testing.T) {
	srv := utils.Setup(Bitcoin)

	resp, err := http.Get(fmt.Sprintf("%s/btc2usd", srv.URL))
	if err != nil {
		t.Fatal(err)
	}

	// Closing response body if there is one.
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected 200, got %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	type Content struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	}

	var Body struct {
		Content `json:"data"`
	}

	err = json.Unmarshal(bodyBytes, &Body)
	if err != nil {
		t.Fatal(err)
	}

	if Body.Content.Base == "" {
		t.Fatalf("Expected base to have content.")
	}

	if Body.Content.Currency == "" {
		t.Fatalf("Expected currency to have content.")
	}

	if Body.Content.Amount == "" {
		t.Fatalf("Expected amount to have content.")
	}

	utils.Teardown(srv)
}
