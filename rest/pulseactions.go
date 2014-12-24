package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"appengine"
	"appengine/user"

	"db"
)

type PulseActions struct {
	RestActions
	Dao db.PulseDao
}

// TODO move boilerplate
func (pa *PulseActions) Get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	parseErr := r.ParseForm()
	if parseErr != nil {
		http.Error(w, parseErr.Error(), http.StatusInternalServerError)
	}

	year, month, day := time.Now().Date()

	pulse := pa.Dao.Get(u.Email, year, month, day)

	pulseJs, err := json.Marshal(pulse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(pulseJs)
}

func (ph *PulseActions) Post(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "POST not implemented", http.StatusInternalServerError)
}

func (ph *PulseActions) Put(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "PUT not implemented", http.StatusInternalServerError)
}

func (ph *PulseActions) Delete(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "DELETE not implemented", http.StatusInternalServerError)
}
