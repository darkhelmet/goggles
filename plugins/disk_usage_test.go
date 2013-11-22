package plugins_test

import (
    "github.com/darkhelmet/goggles/plugins"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestDiskUsage(t *testing.T) {
    Convey("NewDiskUsage", t, func() {
        Convey("requires a path", func() {
            p, err := plugins.NewDiskUsage(plugins.Params{})
            So(p, ShouldBeNil)
            So(err.Error(), ShouldEqual, "DiskUsage requires a Path parameter")
        })

        Convey("requires a string path", func() {
            p, err := plugins.NewDiskUsage(plugins.Params{"Path": 5})
            So(p, ShouldBeNil)
            So(err.Error(), ShouldEqual, "failed decoding params: int doesn't match expected type string for Path")
        })

        Convey("requires a string binary", func() {
            p, err := plugins.NewDiskUsage(plugins.Params{"Path": "/", "Binary": 5})
            So(p, ShouldBeNil)
            So(err.Error(), ShouldEqual, "failed decoding params: int doesn't match expected type string for Binary")
        })
    })
}
