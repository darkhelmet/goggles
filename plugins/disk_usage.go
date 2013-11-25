package plugins

import (
    "errors"
    "fmt"
    i "github.com/darkhelmet/goggles/influxdb"
    "syscall"
)

func init() {
    Register("DiskUsage", NewDiskUsage)
}

// Checks disk usage using df
type DiskUsage struct {
    // The path to examine
    Paths []string
}

func (du *DiskUsage) Run(points chan i.P) error {
    for _, path := range du.Paths {
        var stat syscall.Statfs_t
        if err := syscall.Statfs(path, &stat); err != nil {
            return err
        }

        bsize := uint64(stat.Bsize)
        available := bsize * stat.Bavail / 1024 / 1024
        total := bsize * stat.Blocks / 1024 / 1024
        used := total - available
        percent := float64(used) / float64(total) * 100

        points <- i.P{"name": "DiskUsage", "path": path, "used": used, "available": available, "percent": percent, "total": total}
    }
    return nil
}

// Constructor for DiskUsage plugin
func NewDiskUsage(p Params) (Plugin, error) {
    du := &DiskUsage{}

    err := p.Decode(du)
    if err != nil {
        return nil, fmt.Errorf("failed decoding params: %s", err)
    }

    if len(du.Paths) == 0 {
        return nil, errors.New("DiskUsage requires at least one path to check, Paths is empty")
    }

    return du, nil
}
