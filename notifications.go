package liveliness

import (
	"log"
	"syscall"
)

func IsReady() bool {
	return isReady.Load().(bool)
}

func IsHealthy() bool {
	return isHealthy.Load().(bool)
}

func SignalIsReady() {
	isReady.Store(true)
	isHealthy.Store(true)

	log.Print(isHealthy.Load())
}

func SignalIsNotReady() {
	isReady.Store(false)
}

func SignalShutdown() {
	syscall.SIGTERM.Signal()
}
