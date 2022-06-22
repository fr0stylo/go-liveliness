package liveliness

import (
	"net/http"
	"sync/atomic"
)

type Probe struct {
	probe *atomic.Value
}

func (l *Probe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if condition := l.probe.Load().(bool); condition {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service Unavailable"))
	}
}

func NewProbe(probe *atomic.Value) *Probe {
	return &Probe{
		probe,
	}
}
