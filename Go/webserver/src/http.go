package src

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

var jsonContentType = "application/json"

type httpServer struct {
	People *People
}

func NewHTTPServer(addr string) *http.Server {
	server := &httpServer{
		People: &People{},
	}

	r := &http.ServeMux{}
	r.HandleFunc("GET /pessoas/{id}", server.handleGet)
	r.HandleFunc("POST /pessoas", server.handlePost)
	r.HandleFunc("GET /contagem-pessoas", server.handleCount)

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

type IDDocument struct {
	Id string `json:"id"`
}

func (h *httpServer) handleGet(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")

	person, err := h.People.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(person)
}

func (h *httpServer) handlePost(w http.ResponseWriter, req *http.Request) {
	var newPerson InsertPerson
	err := json.NewDecoder(req.Body).Decode(&newPerson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := h.People.Insert(newPerson)
	res := IDDocument{Id: id}
	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(res)
}

func (h *httpServer) handleCount(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, strconv.Itoa(len(h.People.people)))
}
