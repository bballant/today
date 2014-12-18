package main

import (
	"encoding/json"
	"event"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResonseWriter, r *http.Request) {
	g := Event{
		Date: timeNow(),
	}
}
