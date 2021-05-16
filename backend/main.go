package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dunkbing/sfw-checker-viet/backend/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "backend", log.LstdFlags)
	sm := mux.NewRouter()

	getRoute := sm.Methods(http.MethodGet).Subrouter()
	getRoute.HandleFunc("/posts", handlers.GetAll)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Receive terminate, shutting down", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
