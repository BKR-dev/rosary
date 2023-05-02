package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

func main() {

	t := &tls.Config{
		MinVersion: 12,
	}

	tr := &http.Transport{
		MaxIdleConns:        5,
		MaxIdleConnsPerHost: 5,
		TLSClientConfig:     t,
	}

	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: tr,
	}

	client.CloseIdleConnections()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

}
