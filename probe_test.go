package liveliness

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestProbe_ServeHTTP(t *testing.T) {
	type fields struct {
		AtomicProbe *atomic.Value
	}

	type want struct {
		statusCode int
		body       string
	}
	tests := []struct {
		name   string
		fields fields
		want   want
		prep   func(*fields)
	}{
		{
			name:   "TestProbe_ServeHTTP_OK",
			fields: fields{AtomicProbe: &atomic.Value{}},
			want:   want{statusCode: http.StatusOK, body: "OK"},
			prep: func(f *fields) {
				f.AtomicProbe.Store(true)
			},
		},
		{
			name:   "TestProbe_ServeHTTP_ServiceUnavailable",
			fields: fields{AtomicProbe: &atomic.Value{}},
			want:   want{statusCode: http.StatusServiceUnavailable, body: "Service Unavailable"},
			prep: func(f *fields) {
				f.AtomicProbe.Store(false)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			l := NewProbe(tt.fields.AtomicProbe)
			tt.prep(&tt.fields)

			r, _ := http.NewRequest("GET", "/", nil)
			l.ServeHTTP(rr, r)

			if rr.Code != tt.want.statusCode {
				t.Errorf("Probe.ServeHTTP() statusCode = %v, want %v", rr.Code, tt.want.statusCode)
			}
			if rr.Body.String() != tt.want.body {
				t.Errorf("Probe.ServeHTTP() body = %v, want %v", rr.Body.String(), tt.want.body)
			}
		})
	}
}
