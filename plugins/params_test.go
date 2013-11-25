package plugins_test

import (
    "github.com/darkhelmet/goggles/plugins"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestDecode(t *testing.T) {
    Convey("string", t, func() {
        type S struct {
            S string
        }

        Convey("should work", func() {
            var s S
            p := plugins.Params{"S": "string"}
            err := p.Decode(&s)
            So(err, ShouldBeNil)
            So(s.S, ShouldEqual, "string")
        })

        Convey("should fail", func() {
            var s S
            p := plugins.Params{"S": 5}
            err := p.Decode(&s)
            So(err, ShouldNotBeNil)
            So(err.Error(), ShouldEqual, "int doesn't match expected type string for S")
            So(s.S, ShouldEqual, "")
        })
    })

    Convey("int", t, func() {
        type I struct {
            I int
        }

        Convey("should work", func() {
            var i I
            p := plugins.Params{"I": 5}
            err := p.Decode(&i)
            So(err, ShouldBeNil)
            So(i.I, ShouldEqual, 5)
        })

        Convey("should fail", func() {
            var i I
            p := plugins.Params{"I": "string"}
            err := p.Decode(&i)
            So(err, ShouldNotBeNil)
            So(err.Error(), ShouldEqual, "string doesn't match expected type int for I")
            So(i.I, ShouldEqual, 0)
        })
    })

    Convey("[]string", t, func() {
        type SS struct {
            SS []string
        }

        Convey("should work", func() {
            var ss SS
            p := plugins.Params{"SS": []string{"string"}}
            err := p.Decode(&ss)
            So(err, ShouldBeNil)
            So(ss.SS, ShouldResemble, []string{"string"})
        })

        Convey("should work on []interface{}", func() {
            var ss SS
            p := plugins.Params{"SS": []interface{}{"batman"}}
            err := p.Decode(&ss)
            So(err, ShouldBeNil)
            So(ss.SS, ShouldResemble, []string{"batman"})
        })

        Convey("should fail", func() {
            var ss SS
            p := plugins.Params{"SS": 5}
            err := p.Decode(&ss)
            So(err, ShouldNotBeNil)
            So(err.Error(), ShouldEqual, "int doesn't match expected type []string for SS")
            So(len(ss.SS), ShouldEqual, 0)
        })
    })
}
