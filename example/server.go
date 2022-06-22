package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fr0stylo/go-liveliness"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
		w.WriteHeader(http.StatusOK)
	}))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	livenessServer := liveliness.NewLivelinessServer(liveliness.NewServerOptions())

	livenessServer.RegisterExitHandler(server.Close)

	go livenessServer.Start()

	<-time.After(5 * time.Second)
	liveliness.SignalIsReady()

	log.Fatal(server.ListenAndServe())
}
