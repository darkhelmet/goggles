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
    // The name of the machine.
    Hostname string
    // The InfluxDB config
    InfluxDB influxdb.Config
    // Which checks to perform
    Checks []Check
}

// Load a configuration JSON file from disk
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
