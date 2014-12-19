package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"appengine"
	"appengine/user"
)

func NewEvent(userEmail string, values url.Values) *Event {
	weight, err := strconv.Atoi(values.Get("weight"))
	if err != nil {
		weight = -1
	}

	happiness, err := strconv.Atoi(values.Get("happiness"))
	if err != nil {
		happiness = -1
	}

	drink, err := strconv.Atoi(values.Get("happiness"))
	if err != nil {
		drink = -1
	}

	eat, err := strconv.Atoi(values.Get("eat"))
	if err != nil {
		eat = -1
	}

	exercise, err := strconv.Atoi(values.Get("exercise"))
	if err != nil {
		exercise = -1
	}

	return &Event{
		Email:     userEmail,
		Date:      time.Now(),
		Weight:    weight,
		Happiness: happiness,
		Drink:     drink,
		Eat:       eat,
		Exercise:  exercise,
	}
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
	}

	parseErr := r.ParseForm()
	if parseErr != nil {
		http.Error(w, parseErr.Error(), http.StatusInternalServerError)
	}

	event := NewEvent(u.Email, r.Form)

	fmt.Fprint(w, "Hello, world!\n")
	var jj, _ = json.Marshal(event)
	fmt.Fprintf(w, "%#v\n", event)
	fmt.Fprintf(w, "%v", string(jj))
}
