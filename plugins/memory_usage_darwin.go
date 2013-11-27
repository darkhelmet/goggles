package plugins

/*
#include <unistd.h>
#include <mach/mach_host.h>
#include <mach/host_info.h>
*/
import "C"

import (
    "fmt"
    i "github.com/darkhelmet/goggles/influxdb"
    "unsafe"
)

func (mu *MemoryUsage) Run(points chan i.P) error {
    if err := sysctlByName("vm.swapusage", &mu.Swap); err != nil {
        return err
    }

    var total uint64
    if err := sysctlByName("hw.memsize", &total); err != nil {
        return err
    }

    var vmstat C.vm_statistics_data_t
    if err := host_statistics(&vmstat); err != nil {
        return err
    }

    pagesize := uint64(C.getpagesize())
    inactive := uint64(vmstat.inactive_count) * pagesize

    free := uint64(vmstat.free_count) * pagesize
    used := total - free
    actualFree := free + inactive
    actualUsed := used - inactive

    points <- i.P{
        "name":              "MemoryUsage",
        "swap_percent_used": float64(mu.Swap.Used) / float64(mu.Swap.Total) * 100,
        "swap_total":        mu.Swap.Total,
        "swap_avail":        mu.Swap.Avail,
        "swap_used":         mu.Swap.Used,
        "total":             total,
        "free":              free,
        "used":              used,
        "actual_used":       actualUsed,
        "actual_free":       actualFree,
        "percent_used":      float64(actualUsed) / float64(total) * 100,
    }

    return nil
}

func host_statistics(vmstat *C.vm_statistics_data_t) error {
    var count C.mach_msg_type_number_t = C.HOST_VM_INFO_COUNT

    status := C.host_statistics(
        C.host_t(C.mach_host_self()),
        C.HOST_VM_INFO,
        C.host_info_t(unsafe.Pointer(vmstat)),
        &count)

    if status != C.KERN_SUCCESS {
        return fmt.Errorf("host_statistics=%d", status)
    }

    return nil
}
