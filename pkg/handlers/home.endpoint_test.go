package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/valp0/simple-api/utils"
)

func TestHome(t *testing.T) {
	srv := utils.Setup(Home)

	resp, err := http.Get(fmt.Sprintf("%s/", srv.URL))
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
		Message string `json:"message"`
	}

	var Body struct {
		Content `json:"data"`
	}

	err = json.Unmarshal(bodyBytes, &Body)
	if err != nil {
		t.Fatal(err)
	}

	if Body.Content.Message != "Please use either /hello or /btc2usd endpoints." {
		t.Fatalf("Expected \"Please use either /hello or /btc2usd endpoints.\", got \"%s\"", Body.Content.Message)
	}

	utils.Teardown(srv)
}
