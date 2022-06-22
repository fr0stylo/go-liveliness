package liveliness

import (
	"syscall"
)

// IsReady returns true if the service is ready to serve requests.
func IsReady() bool {
	return isReady.Load().(bool)
}

// IsHealthy returns true if the service is healthy.
func IsHealthy() bool {
	return isHealthy.Load().(bool)
}

// SignalIsReady signals that the service is ready to serve requests.
func SignalIsReady() {
	isReady.Store(true)
	isHealthy.Store(true)
}

// SignalIsNotReady signals that the service is not ready to serve requests.
func SignalIsNotReady() {
	isReady.Store(false)
}

// SignalShutdown signals that the service is shutting down.
func SignalShutdown() {
	syscall.SIGTERM.Signal()
}
