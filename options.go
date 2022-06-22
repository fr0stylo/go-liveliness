package liveliness

import "time"

//ServerOptions defines the options for the server
type ServerOptions struct {
	Addr                   string
	DetectKubernetes       bool
	HealthEndpoint         string
	ReadyEndpoint          string
	GracePeriod            time.Duration
	TerminationGracePeriod time.Duration
}

//NewServerOptions returns a new ServerOptions struct with default values
//	Addr: ":9090"
//	DetectKubernetes: true
//	HealthEndpoint: "/healthz"
//	ReadyEndpoint: "/readyz"
//	GracePeriod: 5s
//	TerminationGracePeriod: 5s
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		Addr:                   ":9090",
		DetectKubernetes:       true,
		HealthEndpoint:         "/healthz",
		ReadyEndpoint:          "/readyz",
		GracePeriod:            time.Second * 5,
		TerminationGracePeriod: time.Second * 5,
	}
}

//LivelinessServer defines the server port address for the liveliness
func (o *ServerOptions) WithAddr(addr string) *ServerOptions {
	o.Addr = addr
	return o
}

//WithDetectKubernetes sets the DetectKubernetes option which controls if the server should detect if it is running in Kubernetes
func (o *ServerOptions) WithDetectKubernetes(detectKubernetes bool) *ServerOptions {
	o.DetectKubernetes = detectKubernetes
	return o
}

//WithHealthEndpoint sets the HealthEndpoint option which controls the endpoint for the health check
func (o *ServerOptions) WithHealthEndpoint(healthEndpoint string) *ServerOptions {
	o.HealthEndpoint = healthEndpoint
	return o
}

//WithReadyEndpoint sets the ReadyEndpoint option which controls the endpoint for the ready check
func (o *ServerOptions) WithReadyEndpoint(readyEndpoint string) *ServerOptions {
	o.ReadyEndpoint = readyEndpoint
	return o
}

//WithGracePeriod sets the GracePeriod option which controls the grace period before shutdown
func (o *ServerOptions) WithGracePeriod(gracePeriod time.Duration) *ServerOptions {
	o.GracePeriod = gracePeriod
	return o
}

func (o *ServerOptions) WithTerminationGracePeriod(terminationGracePeriod time.Duration) *ServerOptions {
	o.TerminationGracePeriod = terminationGracePeriod
	return o
}
