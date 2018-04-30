package nats

import (
	"sync"

	"github.com/elojah/services"
)

// Namespaces maps configs used for nats service with config file namespaces.
type Namespaces struct {
	Nats services.Namespace
}

// Launcher represents a nats launcher.
type Launcher struct {
	*services.Configs
	ns Namespaces

	service *Service
	m       sync.Mutex
}

// NewLauncher returns a new nats Launcher.
func (Service *Service) NewLauncher(ns Namespaces, nsRead ...services.Namespace) *Launcher {
	return &Launcher{
		Configs: services.NewConfigs(nsRead...),
		service: Service,
		ns:      ns,
	}
}

// Up starts the nats service with new configs.
func (l *Launcher) Up(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	cfg := Config{}
	if err := cfg.Dial(configs[l.ns.Nats]); err != nil {
		// Add namespace key when returning error with logrus
		return err
	}
	return l.service.Dial(cfg)
}

// Down stops the nats service.
func (l *Launcher) Down(configs services.Configs) error {
	l.m.Lock()
	defer l.m.Unlock()

	l.service.Close()
	return nil
}
