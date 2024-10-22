package src

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var jsonContentType = "application/json"

// ERRORS
var ErrIDNotFound = fmt.Errorf("ID not found")
var ErrInsertPerson = fmt.Errorf("Error inserting person")
var ErrQueryParamObrigatory = fmt.Errorf("Missing query param 't' in request")
var ErrSearchPeople = fmt.Errorf("Failed to search for people")

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
	r.HandleFunc("GET /pessoas", server.handleSearch)
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
		http.Error(w, ErrInsertPerson.Error(), http.StatusBadRequest)
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

func (h *httpServer) handleSearch(w http.ResponseWriter, req *http.Request) {
	searchParam := req.URL.Query().Get("t")

	if searchParam == "" {
		http.Error(w, ErrQueryParamObrigatory.Error(), http.StatusBadRequest)
		return
	}

	people, err := h.People.Search(h.db, searchParam)
	if err != nil {
		fmt.Printf("Error searching for people %s\n", err)
		http.Error(w, ErrSearchPeople.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", jsonContentType)
	json.NewEncoder(w).Encode(people)
}
