package handlers

import (
	"log"
	"net/http"
)

// Goodbye handler
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye creates new goodbye handler
// and inject a logger into it
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye handler")
	rw.Write([]byte("Goodbye!"))
}
