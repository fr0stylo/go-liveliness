package liveliness

import (
	"net/http"
	"sync/atomic"
)

type HealthzProbe struct {
	probe *atomic.Value
}

func (l *HealthzProbe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func NewHealthzProbe(probe *atomic.Value) *HealthzProbe {
	return &HealthzProbe{
		probe,
	}
}
