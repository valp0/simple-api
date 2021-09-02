package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	type Greeting struct {
		Value string `json:"value"`
	}

	type Body struct {
		Data Greeting `json:"data"`
	}

	w.Header().Set("Content-Type", "application/json")

	// If we send a name query param, it will be used to greet.
	name, ok := r.URL.Query()["name"]
	var resp Body
	var output string

	// Will return "Hello, World." if no name query param is sent.
	if !ok || len(name[0]) < 1 {
		output = "Hello, World."
	} else {
		output = fmt.Sprintf("Hello, %s.", name[0])
	}

	// We create our response body.
	resp.Data = Greeting{output}

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
	}
}
