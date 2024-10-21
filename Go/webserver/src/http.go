package src

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var jsonContentType = "application/json"

type httpServer struct {
	People *People
	db     *DB
}

func NewHTTPServer(addr string, db *DB) *http.Server {
	server := &httpServer{
		People: &People{},
		db:     db,
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

	person, err := h.People.Get(h.db, id)
	if err != nil {
		fmt.Printf("Error getting person %s\n", err)
		http.Error(w, ErrIDNotFound.Error(), http.StatusBadRequest)
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

	id, err := h.People.Insert(h.db, newPerson)
	if err != nil {
		fmt.Printf("Error insert person %s\n", err)
		http.Error(w, "Error inserting person", http.StatusBadRequest)
		return
	}

	res := IDDocument{Id: id}
	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(res)
}

func (h *httpServer) handleCount(w http.ResponseWriter, req *http.Request) {
	val, _ := h.db.CountAllPeople()
	io.WriteString(w, strconv.Itoa(val))
}
