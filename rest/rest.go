package rest

import (
	"net/http"

	"appengine"
	"appengine/user"
)

type RestActions interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type RestHandler struct {
	RestActions
}

func (h *RestHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case "GET":
		h.Get(w, r)
	case "POST":
		h.Post(w, r)
	case "PUT":
		h.Put(w, r)
	case "DELETE":
		h.Delete(w, r)
	default:
		http.Error(w, "Unsupported Action", http.StatusInternalServerError)
	}
}
