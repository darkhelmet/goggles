// Server monitoring in Go
package main

import (
    "flag"
    "github.com/darkhelmet/goggles/config"
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

func main() {
    config, err := config.LoadFile(configPath)
    if err != nil {
        log.Fatalf("failed loading config file (%s): %s", configPath, err)
    }

    if len(config.Checks) == 0 {
        log.Println("no checks, exiting")
        os.Exit(0)
    }

    log.Printf("%v", config)
}
