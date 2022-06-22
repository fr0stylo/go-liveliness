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

type LivelinessServer struct {
	server       http.Server
	shouldRun    bool
	gracePeriod  time.Duration
	exitHandlers []func() error
}

func (l *LivelinessServer) Start() {
	isHealthy.Store(false)
	isReady.Store(false)

	mux := http.NewServeMux()
	mux.Handle("/readyz", NewReadyzProbe(isReady))
	mux.Handle("/healthz", NewHealthzProbe(isHealthy))

	l.server.Handler = mux

	notification := make(chan os.Signal, 1)
	signal.Notify(notification, syscall.SIGTERM, syscall.SIGINT)

	if l.shouldRun {
		go l.server.ListenAndServe()
	}

	<-notification

	SignalIsNotReady()

	<-time.After(l.gracePeriod)

	for _, handler := range l.exitHandlers {
		handler()
	}

	l.Stop()
}

func (l *LivelinessServer) Stop() error {
	if l.shouldRun {
		return l.server.Close()
	}

	return nil
}

func (l *LivelinessServer) SignalIsHealthy() {
	isHealthy.Store(true)
}

func (l *LivelinessServer) RegisterExitHandler(handler ...func() error) {
	l.exitHandlers = append(l.exitHandlers, handler...)
}

func NewLivelinessServer(addr string, detectKubernetes bool, gracePeriod time.Duration) *LivelinessServer {
	server := http.Server{}
	server.Addr = addr

	shouldRun := true
	if detectKubernetes && os.Getenv("KUBERNETES_SERVICE_HOST") != "" {
		shouldRun = false
	}

	return &LivelinessServer{
		server,
		shouldRun,
		gracePeriod,
		[]func() error{},
	}
}
