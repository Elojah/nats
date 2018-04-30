package nats

import (
	"errors"
	"time"
)

// Config is a Stream server config.
type Config struct {
	Address       string
	MaxReconnect  int
	ReconnectWait time.Duration
	Timeout       time.Duration
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c == rhs
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	var err error
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cAddress, ok := fconf["address"]
	if !ok {
		return errors.New("missing key address")
	}
	if c.Address, ok = cAddress.(string); !ok {
		return errors.New("key address invalid. must be string")
	}

	cReconnectWait, ok := fconf["reconnect_wait"]
	if !ok {
		return errors.New("missing key reconnect_wait")
	}
	cReconnectWaitString, ok := cReconnectWait.(string)
	if !ok {
		return errors.New("key reconnect_wait invalid. must be string")
	}
	if c.ReconnectWait, err = time.ParseDuration(cReconnectWaitString); err != nil {
		return err
	}

	cTimeout, ok := fconf["timeout"]
	if !ok {
		return errors.New("missing key timeout")
	}
	cTimeoutString, ok := cTimeout.(string)
	if !ok {
		return errors.New("key timeout invalid. must be string")
	}
	if c.Timeout, err = time.ParseDuration(cTimeoutString); err != nil {
		return err
	}

	return nil
}
