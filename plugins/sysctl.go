package plugins

import (
    "bytes"
    "encoding/binary"
    "log"
    "syscall"
    "unsafe"
)

// From https://github.com/cloudfoundry/gosigar/blob/master/sigar_darwin.go
// Copyright (c) 2012 VMware, Inc.
func sysctlByName(name string, data interface{}) (err error) {
    val, err := syscall.Sysctl(name)
    if err != nil {
        log.Printf("%s", err)
        return err
    }

    buf := []byte(val)

    switch v := data.(type) {
    case *uint64:
        *v = *(*uint64)(unsafe.Pointer(&buf[0]))
        return
    }

    bbuf := bytes.NewBuffer([]byte(val))
    return binary.Read(bbuf, binary.LittleEndian, data)
}
