# simple-api

This is just a simple REST API that supports the following endpoints:
  - /hello -> This endpoint accepts a "name" query param and will respond with "Hello, {name}." if a name param is specified, or with "Hello, World." if no name param is specified.
  - /btc2usd -> This endpoint doesn't accept any query params. It will respond with the current BTC-to-USD exchange rate. It uses the [coinbase API](https://developers.coinbase.com/).
  - / -> I only added this endpoint to send a JSON message to specify which endpoints are available, so it is just for general information purposes, in case it gets accessed by accident.

This API was built using Go 1.15, to run it you only have to clone or fork this repo, then build the main API file by running ```go build cmd/api/main.go``` and running the generated executable file.

Since no external libraries were used to create this simple API, there is no need to install anything else in order to run it.
