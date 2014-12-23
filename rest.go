package main

import (
	"net/http"
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

type PulseActions struct {
	RestActions
}

func (ph *PulseActions) Get(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "GET not implemented", http.StatusInternalServerError)
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
