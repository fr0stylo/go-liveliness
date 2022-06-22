package liveliness

import (
	"net/http"
	"sync/atomic"
)

type ReadyzProbe struct {
	probe *atomic.Value
}

func (l *ReadyzProbe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if condition := l.probe.Load().(bool); condition {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service Unavailable"))
	}
}

func NewReadyzProbe(probe *atomic.Value) *ReadyzProbe {
	return &ReadyzProbe{
		probe,
	}
}
