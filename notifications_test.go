package liveliness

import (
	"testing"
)

func TestIsReady(t *testing.T) {
	tests := []struct {
		name string
		want bool
		prep func()
	}{
		{"TestIsReady", true, func() { SignalIsReady() }},
		{"TestIsNotReady", false, func() { SignalIsNotReady() }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prep()
			if got := IsReady(); got != tt.want {
				t.Errorf("IsReady() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHealthy(t *testing.T) {
	tests := []struct {
		name string
		want bool
		prep func()
	}{
		{"TestIsHealthy", true, func() { SignalIsReady() }},
		{"TestIsHealthyNotReady", true, func() { SignalIsNotReady() }},
		{"TestIsNotHealthy", false, func() { isHealthy.Store(false) }},
	}
	for _, tt := range tests {
		tt.prep()
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHealthy(); got != tt.want {
				t.Errorf("IsHealthy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignalIsReady(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestSignalIsReady"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SignalIsReady()
			if !IsReady() {
				t.Errorf("SignalIsReady() did not set IsReady() to true")
			}
			if !IsHealthy() {
				t.Errorf("SignalIsReady() did not set IsHealthy() to true")
			}
		})
	}
}

func TestSignalIsNotReady(t *testing.T) {
	tests := []struct {
		name  string
		want  bool
		want2 bool
		prep  func()
	}{
		{"AfterTestSignalIsReady", false, true, func() { SignalIsReady() }},
	}
	for _, tt := range tests {
		tt.prep()
		t.Run(tt.name, func(t *testing.T) {
			SignalIsNotReady()
			if IsReady() {
				t.Errorf("SignalIsNotReady() did not set IsReady() to false")
			}
			if !IsHealthy() {
				t.Errorf("SignalIsNotReady() set IsHealthy() to false")
			}
		})
	}
}
