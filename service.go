package nats

import (
	"github.com/nats-io/go-nats"
)

// Service wraps an NatsStream connection.
type Service struct {
	*nats.Conn
}

// Dial init the Stream server.
func (s *Service) Dial(c Config) error {
	var err error
	opts := nats.Options{
		Url:            c.Address,
		AllowReconnect: true,
		MaxReconnect:   c.MaxReconnect,
		ReconnectWait:  c.ReconnectWait,
		Timeout:        c.Timeout,
	}
	s.Conn, err = opts.Connect()
	return err
}

// Healthcheck returns if database responds.
func (s *Service) Healthcheck() error {
	return nil
}
