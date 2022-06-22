package liveliness

import (
	"log"
	"net/http"
	"sync/atomic"
)

//Probe is a struct that holds an atomic value which controls the state of the probe
type Probe struct {
	probe *atomic.Value
}

//ServeHTTP is the http.HandlerFunc for the probe
// This function will return a 200 if the probe is ready, otherwise it will return a 503
func (l *Probe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print(l.probe.Load())
	if condition := l.probe.Load().(bool); condition {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service Unavailable"))
	}
}

//NewProbe creates a new probe with the given atomic value
func NewProbe(probe *atomic.Value) *Probe {
	return &Probe{
		probe,
	}
}
