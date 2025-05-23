package handler

import (
	"encoding/json"
	"github.com/anger-aa/quotes/internal/model"
	"github.com/anger-aa/quotes/internal/storage"
	"github.com/anger-aa/quotes/pkg/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	storage storage.IStorage
}

func NewHandler(storage storage.IStorage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) AddQuote(w http.ResponseWriter, r *http.Request) {
	var req model.Quote

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input parameters", http.StatusBadRequest)
		return
	}

	quote := h.storage.AddQuote(req)

	err = response.JSON(w, "Quote successfully created", quote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	quotes, err := h.storage.GetAllQuotes(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = response.JSON(w, "Quotes", quotes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.storage.GetRandomQuote()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = response.JSON(w, "Random quote", quote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.storage.DeleteQuote(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = response.JSON(w, "Successfully deleted", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
