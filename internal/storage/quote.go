package storage

import (
	"errors"
	"fmt"
	"github.com/anger-aa/quotes/internal/model"
	"math/rand"
	"sync"
)

type Storage struct {
	storage map[int]*model.Quote
	quoteID int
	mu      sync.Mutex
}

type IStorage interface {
	AddQuote(quote model.Quote) *model.Quote
	GetAllQuotes(author string) ([]*model.Quote, error)
	GetRandomQuote() (*model.Quote, error)
	DeleteQuote(id int) error
}

func NewStorage() *Storage {
	return &Storage{
		storage: make(map[int]*model.Quote),
		quoteID: 1,
	}
}

func (s *Storage) AddQuote(quote model.Quote) *model.Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	quote.ID = s.quoteID
	s.storage[s.quoteID] = &quote
	s.quoteID++

	return &quote
}

func (s *Storage) GetAllQuotes(author string) ([]*model.Quote, error) {
	var quotes []*model.Quote

	s.mu.Lock()
	defer s.mu.Unlock()

	for id, quote := range s.storage {
		if author == "" || author == quote.Author {
			quotes = append(quotes, &model.Quote{
				ID:     id,
				Author: quote.Author,
				Quote:  quote.Quote,
			})
		}
	}

	if author != "" && len(quotes) == 0 {
		return nil, fmt.Errorf("could not find quotes for author %s", author)
	}

	return quotes, nil
}

func (s *Storage) GetRandomQuote() (*model.Quote, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.storage) == 0 {
		return nil, errors.New("there is no quotes")
	}

	var keys []int
	for id := range s.storage {
		keys = append(keys, id)
	}

	i := rand.Intn(len(keys))

	return s.storage[keys[i]], nil
}

func (s *Storage) DeleteQuote(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.storage[id]; !ok {
		return fmt.Errorf("quote with id %d not found", id)
	}
	delete(s.storage, id)

	return nil
}
