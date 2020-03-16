package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/orkhan-huseyn/microservices-w-go/data"
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
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
	rw.Write(d)
}
