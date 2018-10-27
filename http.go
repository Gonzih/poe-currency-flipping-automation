package main

import (
	"io"
	"net/http"
)

func GET(url string) (io.ReadCloser, error) {
	// log.Printf("Executing GET %s", url)
	response, err := http.Get(url)
	return response.Body, err
}
