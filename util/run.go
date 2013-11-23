package util

import (
    "bufio"
    "bytes"
    "os/exec"
)

func Run(name string, args ...string) ([]byte, error) {
    cmd := exec.Command(name, args...)
    return cmd.CombinedOutput()
}

func Lines(data []byte) []string {
    scanner := bufio.NewScanner(bytes.NewReader(data))
    lines := make([]string, 0)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}
