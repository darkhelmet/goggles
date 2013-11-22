package config

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/darkhelmet/goggles/influxdb"
    "os"
)

type Check struct {
    Plugin string
}

type Config struct {
    Hostname string
    InfluxDB influxdb.Config
    Checks   []Check
}

func LoadFile(path string) (*Config, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    var config Config
    err = decoder.Decode(&config)
    if err != nil {
        return nil, err
    }

    if config.Hostname == "" {
        return nil, errors.New("missing Hostname")
    }

    if err = config.InfluxDB.SetDefaultsAndVerify(); err != nil {
        return nil, fmt.Errorf("invalid InfluxDB config: %s", err)
    }

    return &config, nil
}
