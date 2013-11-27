package plugins

func init() {
    Register("MemoryUsage", NewMemoryUsage)
}

type SwapUsage struct {
    Total, Avail, Used uint64
}

type MemoryUsage struct {
    Swap SwapUsage
}

func NewMemoryUsage(p Params) (Plugin, error) {
    return new(MemoryUsage), nil
}
