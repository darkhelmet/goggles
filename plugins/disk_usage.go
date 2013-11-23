package plugins

import (
    "errors"
    "fmt"
    i "github.com/darkhelmet/goggles/influxdb"
    . "github.com/darkhelmet/goggles/util"
)

func init() {
    Register("DiskUsage", NewDiskUsage)
}

// Checks disk usage using df
type DiskUsage struct {
    // The path to examine
    Path string
    // The name of or path to the df binary (default: df)
    Binary string
}

func (du *DiskUsage) Run(points chan i.P) error {
    data, err := Run(du.Binary, "-m", du.Path)
    if err != nil {
        return err
    }

    lines := Lines(data)
    if len(lines) != 2 {
        return fmt.Errorf("expected 2 lines, got %d", len(lines))
    }

    var filesystem string
    var blocks, used, available, percent uint
    _, err = fmt.Sscanf(lines[1], "%s%d%d%d%d%%", &filesystem, &blocks, &used, &available, &percent)
    if err != nil {
        return fmt.Errorf("failed scanning: %s", err)
    }

    points <- i.P{"name": "DiskUsage", "path": du.Path, "used": used, "available": available, "percent": percent}
    return nil
}

// Constructor for DiskUsage plugin
func NewDiskUsage(p Params) (Plugin, error) {
    du := &DiskUsage{Binary: "df"}

    err := p.Decode(du)
    if err != nil {
        return nil, fmt.Errorf("failed decoding params: %s", err)
    }

    if du.Path == "" {
        return nil, errors.New("DiskUsage requires a Path parameter")
    }

    return du, nil
}
