package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/valp0/simple-api/utils"
)

func TestHelloName(t *testing.T) {
	srv := utils.Setup(Hello)

	resp, err := http.Get(fmt.Sprintf("%s/hello?name=Peter", srv.URL))
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

	type Greeting struct {
		Value string `json:"value"`
	}

	var Body struct {
		Data Greeting `json:"data"`
	}

	err = json.Unmarshal(bodyBytes, &Body)
	if err != nil {
		t.Fatal(err)
	}

	if Body.Data.Value != "Hello, Peter." {
		t.Fatalf("Expected \"Hello, Peter.\", got \"%s\"", Body.Data.Value)
	}

	utils.Teardown(srv)
}

func TestHelloWorld(t *testing.T) {
	srv := utils.Setup(Hello)

	resp, err := http.Get(fmt.Sprintf("%s/hello", srv.URL))
	if err != nil {
		t.Fatal(err)
	}

	// Closing response body if there is one.
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected 200, got %d", resp.StatusCode)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	type Greeting struct {
		Value string `json:"value"`
	}

	var Body struct {
		Data Greeting `json:"data"`
	}

	err = json.Unmarshal(bodyBytes, &Body)
	if err != nil {
		t.Fatal(err)
	}

	if Body.Data.Value != "Hello, World." {
		t.Fatalf("Expected \"Hello, World.\", got \"%s\"", Body.Data.Value)
	}

	utils.Teardown(srv)
}
