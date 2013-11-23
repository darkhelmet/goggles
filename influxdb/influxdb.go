package influxdb

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "time"
)

type P map[string]interface{}

func (p P) String(key string) (string, error) {
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

func (p P) Delete(key string) {
    delete(p, key)
}

func (p P) Set(key string, value interface{}) {
    p[key] = value
}

type InfluxDB struct {
    Config
}

type Series struct {
    Name    string          `json:"name"`
    Columns []string        `json:"columns"`
    Points  [][]interface{} `json:"points"`
}

func (s *Series) HasColumns(name string) bool {
    for _, column := range s.Columns {
        if column == name {
            return true
        }
    }
    return false
}

func (s *Series) Append(p P) {
    for key := range p {
        if !s.HasColumns(key) {
            s.Columns = append(s.Columns, key)
        }
    }

    points := make([]interface{}, 0)
    for _, key := range s.Columns {
        points = append(points, p[key])
    }
    s.Points = append(s.Points, points)
}

func (i *InfluxDB) URL() string {
    query := make(url.Values)
    query.Add("u", i.Username)
    query.Add("p", i.Password)
    query.Add("time_precision", "s")
    return fmt.Sprintf("http://%s:%d/db/%s/series?%s", i.Host, i.Port, i.Database, query.Encode())
}

func (i *InfluxDB) Report(reports <-chan interface{}) error {
    data := make(map[string]*Series)
    for report := range reports {
        p := report.(P)

        name, err := p.String("name")
        if err != nil {
            log.Println(err)
            continue
        }

        p.Delete("name")
        p.Set("time", time.Now().Unix())

        series, ok := data[name]
        if !ok {
            series = &Series{
                Name:    name,
                Columns: []string{"time"},
            }
            data[name] = series
        }

        series.Append(p)
    }

    payload := make([]*Series, 0, len(data))
    for _, series := range data {
        log.Printf("%v", series)
        payload = append(payload, series)
    }
    body, err := json.Marshal(payload)
    if err != nil {
        return fmt.Errorf("failed to marshal payload: %s", err)
    }

    resp, err := http.Post(i.URL(), "application/json", bytes.NewReader(body))
    if err != nil {
        return fmt.Errorf("failed to send data: %s", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return fmt.Errorf("expected a 200 from InfluxDB, got %d", resp.StatusCode)
    }

    return nil
}
