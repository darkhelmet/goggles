package plugins_test

import (
    "github.com/darkhelmet/goggles/plugins"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestDiskUsage(t *testing.T) {
    Convey("NewDiskUsage", t, func() {
        Convey("require a path", func() {
            p, err := plugins.NewDiskUsage(plugins.Params{})
            So(p, ShouldBeNil)
            So(err.Error(), ShouldEqual, "DiskUsage requires a Path parameter")
        })

        Convey("require a string path", func() {
            p, err := plugins.NewDiskUsage(plugins.Params{"Path": 5})
            So(p, ShouldBeNil)
            So(err.Error(), ShouldEqual, "failed decoding params: int doesn't match expected type string for Path")
        })

        Convey("require a string binary", func() {
            p, err := plugins.NewDiskUsage(plugins.Params{"Path": "/", "Binary": 5})
            So(p, ShouldBeNil)
            So(err.Error(), ShouldEqual, "failed decoding params: int doesn't match expected type string for Binary")
        })
    })

    // Convey("Build", t, func() {
    //     Convey("should return an error if plugin not found", func() {
    //         p, err := plugins.Build("not_found", plugins.Params{"foo": "bar"})
    //         So(p, ShouldBeNil)
    //         So(err.Error(), ShouldEqual, "not_found is not registered")
    //     })
    // })
}
