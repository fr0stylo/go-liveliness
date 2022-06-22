package liveliness

import (
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

var (
	isReady   = &atomic.Value{}
	isHealthy = &atomic.Value{}
)

// LivelinessServer is a struct containing the http.Server which is used to serve the probe endpoints.
type LivelinessServer struct {
	server       http.Server
	shouldRun    bool
	opts         *ServerOptions
	exitHandlers []func() error
}

// Start starts the liveliness server and registers the probe endpoints.
func (l *LivelinessServer) Start() {
	mux := http.NewServeMux()

	mux.Handle(l.opts.ReadyEndpoint, NewProbe(isReady))
	mux.Handle(l.opts.HealthEndpoint, NewProbe(isHealthy))

	l.server.Handler = mux

	notification := make(chan os.Signal, 1)
	signal.Notify(notification, syscall.SIGTERM, syscall.SIGINT)

	if l.shouldRun {
		go l.server.ListenAndServe()
	}

	<-notification

	SignalIsNotReady()

	<-time.After(l.opts.GracePeriod)

	for _, handler := range l.exitHandlers {
		handler()
	}

	l.Stop()
}

// Stop stops the liveliness server.
func (l *LivelinessServer) Stop() error {
	if l.shouldRun {
		return l.server.Close()
	}

	return nil
}

// RegisterExitHandler registers a function to be called when the server is stopped.
func (l *LivelinessServer) RegisterExitHandler(handler ...func() error) {
	l.exitHandlers = append(l.exitHandlers, handler...)
}

// NewLivelinessServer creates a new liveliness server with the given options.
func NewLivelinessServer(opts *ServerOptions) *LivelinessServer {
	isHealthy.Store(false)
	isReady.Store(false)

	server := http.Server{}
	server.Addr = opts.Addr

	shouldRun := true
	if opts.DetectKubernetes && os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		shouldRun = false
	}

	return &LivelinessServer{
		server,
		shouldRun,
		opts,
		[]func() error{},
	}
}
