package main

import (
	"log"
	"strconv"
)

func s2i(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Error parsing %s to int", s)
	}

	return i
}

func s2f(s string) float64 {
	i, err := strconv.ParseFloat(s, 10)
	if err != nil {
		log.Fatalf("Error parsing %s to float", s)
	}

	return i
}
