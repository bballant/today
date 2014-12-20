package main

import (
	"encoding/json"
	//"fmt"
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/new", new)
}

//curl -v -b dev_appserver_login=test@example.com:False:185804764220139124118 "http://localhost:8080/new?weight=185&happiness=5&drink=4&eat=4&exercise=0"
func new(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	parseErr := r.ParseForm()
	if parseErr != nil {
		http.Error(w, parseErr.Error(), http.StatusInternalServerError)
	}

	pulse := NewPulse(u.Email, r.Form)

	pulseJs, err := json.Marshal(pulse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(pulseJs)

	// fmt.Fprint(w, "Hello, world!\n")
	// var jj, _ = json.Marshal(pulse)
	// fmt.Fprintf(w, "%#v\n", pulse)
	// fmt.Fprintf(w, "%v", string(jj))
}
