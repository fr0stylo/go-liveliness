package liveliness

import (
	"net/http"
	"sync/atomic"
)

type HealthzProbe struct {
	probe *atomic.Value
}

func (l *HealthzProbe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if condition := l.probe.Load().(bool); condition {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service Unavailable"))
	}
}

func NewHealthzProbe(probe *atomic.Value) *HealthzProbe {
	return &HealthzProbe{
		probe,
	}
}
