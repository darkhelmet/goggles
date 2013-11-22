package plugins

import (
    "errors"
    "fmt"
)

func init() {
    Register("DiskUsage", NewDiskUsage)
}

// Checks disk usage using df
type DiskUsage struct {
    Path string
    // The path to the df binary (default: /bin/df)
    Binary string
}

func (du *DiskUsage) Run() error {
    return nil
}

// Constructor for DiskUsage plugin
func NewDiskUsage(p Params) (Plugin, error) {
    du := &DiskUsage{Binary: "/bin/df"}

    err := p.Decode(du)
    if err != nil {
        return nil, fmt.Errorf("failed decoding params: %s", err)
    }

    if du.Path == "" {
        return nil, errors.New("DiskUsage requires a Path parameter")
    }

    return du, nil
}
