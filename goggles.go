// Server monitoring in Go
package main

import (
    "flag"
    "github.com/darkhelmet/goctopus"
    "github.com/darkhelmet/goggles/config"
    "github.com/darkhelmet/goggles/influxdb"
    "github.com/darkhelmet/goggles/plugins"
    "log"
    "os"
)

var (
    configPath string
)

func init() {
    flag.StringVar(&configPath, "config", "goggles.json", "the config file to load")
    flag.Parse()
}

func RunPlugin(name string, p plugins.Plugin) chan influxdb.P {
    ch := make(chan influxdb.P, 1)
    go func() {
        defer close(ch)
        if err := p.Run(ch); err != nil {
            log.Printf("failed to run %s: %s", name, err)
        }
    }()
    return ch
}

func LoadConfigFile() *config.Config {
    config, err := config.LoadFile(configPath)
    if err != nil {
        log.Fatalf("failed loading config file (%s): %s", configPath, err)
    }

    if len(config.Checks) == 0 {
        log.Println("no checks, exiting")
        os.Exit(0)
    }
    return config
}

func main() {
    config := LoadConfigFile()

    channels := make([]interface{}, 0)
    for _, check := range config.Checks {
        p, err := plugins.Build(check.Plugin, check.Params)
        if err != nil {
            log.Printf("failed building plugin: %s", err)
            continue
        }
        channels = append(channels, RunPlugin(check.Plugin, p))
    }
    reports := goctopus.New(channels...).Run()

    db := influxdb.InfluxDB{config.InfluxDB}
    err := db.Report(reports)
    if err != nil {
        log.Printf("reporting failed: %s", err)
    }
}
