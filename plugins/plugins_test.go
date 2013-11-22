package plugins_test

import (
    "github.com/darkhelmet/goggles/plugins"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func Panic(p plugins.Params) (plugins.Plugin, error) {
    return nil, nil
}

func TestRegister(t *testing.T) {
    Convey("Register", t, func() {
        Convey("should panic if registering things twice", func() {
            So(func() {
                plugins.Register("panic", Panic)
                plugins.Register("panic", Panic)
            }, ShouldPanic) // With, errors.New("panic is already registered"))
        })
    })

    Convey("Build", t, func() {
        Convey("should return an error if plugin not found", func() {
            p, err := plugins.Build("not_found", plugins.Params{"foo": "bar"})
            So(p, ShouldBeNil)
            So(err.Error(), ShouldEqual, "not_found is not registered")
        })
    })
}
