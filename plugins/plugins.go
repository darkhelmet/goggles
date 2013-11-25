// Where all the plugins and supporting functions live
package plugins

import (
    "errors"
    "fmt"
    "github.com/darkhelmet/goggles/influxdb"
)

var (
    ParamMissing   = errors.New("param missing")
    ParamWrongType = errors.New("param wrong type")
)

type Plugin interface {
    Run(chan influxdb.P) error
}

type Constructor func(Params) (Plugin, error)

var registry = make(map[string]Constructor)

// Register a plugin for use, panics if plugin already registered
func Register(name string, c Constructor) {
    if _, ok := registry[name]; ok {
        panic(fmt.Errorf("%s is already registered", name))
    }
    registry[name] = c
}

// Builds a plugin from the registry given the name and params
func Build(name string, m Params) (Plugin, error) {
    c, ok := registry[name]
    if ok {
        return c(m)
    }
    return nil, fmt.Errorf("%s is not registered", name)
}
