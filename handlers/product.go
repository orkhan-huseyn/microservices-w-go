package handlers

import (
	"github.com/orkhan-huseyn/microservices-w-go/data"
	"log"
	"net/http"
)

// Products handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a new products handler
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: // handle GET methods
		p.getProducts(rw, r)
	case http.MethodPost: // handle POST methods
		p.getProducts(rw, r)
	default: // handle all other methods that are not implemented
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, _ *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}
