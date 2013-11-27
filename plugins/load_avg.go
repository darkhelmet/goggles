package plugins

/*
#include <stdlib.h>
*/
import "C"

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

func getloadavg() (float64, float64, float64) {
    avg := []C.double{0, 0, 0}

    C.getloadavg(&avg[0], C.int(len(avg)))

    one := float64(avg[0])
    five := float64(avg[1])
    fifteen := float64(avg[2])

    return one, five, fifteen
}

func NewLoadAvg(p Params) (Plugin, error) {
    return new(LoadAvg), nil
}
