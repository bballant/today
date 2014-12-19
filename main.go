package main

import (
	"encoding/json"
	"net/http"
	//"strconv"
	"fmt"
	"time"

	"appengine"
	"appengine/user"
)

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

	g := Event{
		Date: time.Now(),
	}
	fmt.Fprint(w, "Hello, world!\n")
	var jj, _ = json.Marshal(g)
	fmt.Fprintf(w, "%#v\n", g)
	fmt.Fprintf(w, "%v", string(jj))
}
