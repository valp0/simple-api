package main

import (
	"net/http"
	"testing"

	"github.com/valp0/simple-api/pkg/handlers"
	"github.com/valp0/simple-api/utils"
)

func TestMain(t *testing.T) {
	srv := utils.Setup(handlers.Home)

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected 200, got %d", resp.StatusCode)
	}

	utils.Teardown(srv)
}
