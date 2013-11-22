package influxdb

import (
    "errors"
)

const (
    DefaultPort     = 8086
    DefaultDatabase = "goggles"
)

type Config struct {
    Host               string
    Port               int
    Database           string
    Username, Password string
}

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
