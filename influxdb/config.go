package influxdb

import (
    "errors"
)

const (
    DefaultPort     = 8086
    DefaultDatabase = "goggles"
)

type Config struct {
    // The hostname or IP of the InfluxDB server
    Host string
    // The port to connect on (default: 8086)
    Port int
    // The database to use (default: goggles)
    Database string
    // The auth information
    Username, Password string
}

// Deals with required values and setting defaults
func (c *Config) SetDefaultsAndVerify() error {
    if c.Host == "" {
        return errors.New("missing host")
    }

    if c.Username == "" {
        return errors.New("missing username")
    }

    if c.Password == "" {
        return errors.New("missing password")
    }

    if c.Port == 0 {
        c.Port = DefaultPort
    }

    if c.Database == "" {
        c.Database = DefaultDatabase
    }

    return nil
}
