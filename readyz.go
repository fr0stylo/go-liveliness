package liveliness

import (
	"net/http"
	"sync/atomic"
)

type ReadyzProbe struct {
	probe *atomic.Value
}

func (l *ReadyzProbe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func NewReadyzProbe(probe *atomic.Value) *ReadyzProbe {
	return &ReadyzProbe{
		probe,
	}
}
