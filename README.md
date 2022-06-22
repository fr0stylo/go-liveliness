# Go Liveliness

Simple HTTP liveliness probes service to use inside kubernetes environment.

## Description

Simple http service for providing liveliness and readyness probes. Its small and portable package written in pure go without any external dependencies. 

Package provides new http service over dedicated port with two endpoints `/healthz` and `/readyz`.

- Readyness check can be controller using provided notifivation methods - `SignalIsReady` and `SignalIsNotReady`. If service is ready it will return 200 Ok response, otherwise 503 Service Unavailable.
- Liveliness check is set to true whenever signal to service provided as system is ready.


Example usage:

```
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

	livenessServer := liveliness.NewLivelinessServer(":9090", false, time.Second*5)

	livenessServer.RegisterExitHandler(server.Close)

	go livenessServer.Start()

	<-time.After(5 * time.Second)
	liveliness.SignalIsReady()

	log.Fatal(server.ListenAndServe())
}

```

## Getting Started

### Dependencies

* Go 1.18

## Authors

Contributors names and contact info

* Zymantas Maumevicius (fr0stylo)

## Version History

* 0.1
    * Initial Release

## License

This project is licensed under the [NAME HERE] License - see the LICENSE.md file for details
