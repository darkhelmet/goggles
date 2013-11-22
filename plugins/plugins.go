package plugins

import (
    "errors"
    "fmt"
    "reflect"
)

var (
    ParamMissing   = errors.New("param missing")
    ParamWrongType = errors.New("param wrong type")
)

type Plugin interface {
    Run() error
}

type Params map[string]interface{}

func (p Params) Decode(dst interface{}) error {
    v := reflect.ValueOf(dst)
    if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
        return errors.New("dst must be a pointer to struct")
    }
    v = v.Elem()

    for name, value := range p {
        field := v.FieldByName(name)
        if !field.CanSet() {
            return fmt.Errorf("can't set %s", name)
        }
        ft := field.Type()
        switch ft.Kind() {
        case reflect.String:
            s, ok := value.(string)
            if !ok {
                return fmt.Errorf("%T doesn't match expected type %s for %s", value, ft.String(), name)
            }
            field.SetString(s)
        }
    }
    return nil
}

func (p Params) String(key string) (string, error) {
    v, ok := p[key]
    if !ok {
        return "", ParamMissing
    }

    s, ok := v.(string)
    if !ok {
        return "", ParamWrongType
    }

    return s, nil
}

type Constructor func(Params) (Plugin, error)

var registry = make(map[string]Constructor)

func Register(name string, c Constructor) {
    if _, ok := registry[name]; ok {
        panic(fmt.Errorf("%s is already registered", name))
    }
    registry[name] = c
}

func Build(name string, m Params) (Plugin, error) {
    c, ok := registry[name]
    if ok {
        return c(m)
    }
    return nil, fmt.Errorf("%s is not registered", name)
}
