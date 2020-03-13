package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new HelloHandler
// and inject a logger there
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello handler!")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops!", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello, %s\n", d)
}
