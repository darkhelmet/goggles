package plugins

import (
    i "github.com/darkhelmet/goggles/influxdb"
    "runtime"
)

func init() {
    Register("LoadAvg", NewLoadAvg)
}

type LoadAvg struct{}

func (la *LoadAvg) Run(points chan i.P) error {
    one, five, fifteen := getloadavg()

    // Normalize
    nCPUs := runtime.NumCPU()
    one /= float64(nCPUs)
    five /= float64(nCPUs)
    fifteen /= float64(nCPUs)

    points <- i.P{"name": "LoadAvg", "one": one, "five": five, "fifteen": fifteen, "cpus": nCPUs}

    return nil
}

func NewLoadAvg(p Params) (Plugin, error) {
    return new(LoadAvg), nil
}
