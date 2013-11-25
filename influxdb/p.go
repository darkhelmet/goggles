package influxdb

import (
    "fmt"
    "strings"
)

type P map[string]interface{}

func (p P) GetString(key string) (string, error) {
    v, ok := p[key]
    if !ok {
        return "", fmt.Errorf("influxdb: missing %s param", key)
    }
    s, ok := v.(string)
    if !ok {
        return "", fmt.Errorf("influxdb: expected %s to be a string, but was a %T", key, v)
    }
    return s, nil
}

func (p P) String() string {
    parts := make([]string, 0, len(p))
    for key, value := range p {
        parts = append(parts, fmt.Sprintf("%s=%v", key, value))
    }
    return strings.Join(parts, " ")
}

func (p P) Delete(key string) {
    delete(p, key)
}

func (p P) Set(key string, value interface{}) {
    p[key] = value
}
