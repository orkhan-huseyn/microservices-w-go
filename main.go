package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/orkhan-huseyn/microservices-w-go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// initialize handlers
	// hello handler
	hh := handlers.NewHello(l)
	// and the goodbye handler
	gh := handlers.NewGoodbye(l)

	// creating a new serve mux
	// instead of default serve mux
	sm := http.NewServeMux()
	// mapping handlers
	// with the serve mux
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// create our own server
	// specify read, write, idle timeout
	// and the serve mux
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	// making ListenAndServe call
	// non-blocking by calling it
	// inside a go routine
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// create a channel of type os.Signal
	sigChan := make(chan os.Signal)
	// signal.Notify will notify the channel
	// on interrupt (e.g. Ctrl + C) and kill the process
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// this will make the main go routine
	// wait for an interrupt, kill signal
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// make sure we shut down the server gracefully
	// any alive connections will be tolorated up intil 30 seconds
	// and then the server will be shutdown after everyone is done
	tc, cf := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	defer cf()
}
