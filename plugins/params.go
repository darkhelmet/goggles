package plugins

import (
    "errors"
    "fmt"
    "reflect"
)

type Params map[string]interface{}

func NewParamsDecodeError(value interface{}, ft string, name string) error {
    return fmt.Errorf("%T doesn't match expected type %s for %s", value, ft, name)
}

// Decode the params into a struct
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
                return NewParamsDecodeError(value, ft.String(), name)
            }
            field.SetString(s)
        case reflect.Slice:
            vslice := reflect.ValueOf(value)
            if vslice.Kind() != reflect.Slice {
                return NewParamsDecodeError(value, ft.String(), name)
            }

            switch ft.Elem().Kind() {
            case reflect.String:
                length := vslice.Len()
                ss := make([]string, 0, length)
                for i := 0; i < length; i++ {
                    i := vslice.Index(i).Interface()
                    s, ok := i.(string)
                    if !ok {
                        return NewParamsDecodeError(value, ft.String(), name)
                    }
                    ss = append(ss, s)
                }
                field.Set(reflect.ValueOf(ss))
            default:
                return NewParamsDecodeError(value, ft.String(), name)
            }
        case reflect.Int:
            i, ok := value.(int)
            if !ok {
                return NewParamsDecodeError(value, ft.String(), name)
            }
            field.SetInt(int64(i))
        default:
            return fmt.Errorf("don't know how to decode %T into %s", value, ft.String())
        }
    }
    return nil
}
