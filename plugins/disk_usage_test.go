package plugins_test

import (
    "github.com/darkhelmet/goggles/plugins"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestDiskUsage(t *testing.T) {
    Convey("NewDiskUsage", t, func() {
        Convey("requires at least one path", func() {
            p, err := plugins.NewDiskUsage(plugins.Params{})
            So(p, ShouldBeNil)
            So(err.Error(), ShouldEqual, "DiskUsage requires at least one path to check, Paths is empty")
        })

        Convey("requires strings for paths", func() {
            p, err := plugins.NewDiskUsage(plugins.Params{"Paths": []int{5}})
            So(p, ShouldBeNil)
            So(err.Error(), ShouldEqual, "failed decoding params: []int doesn't match expected type []string for Paths")
        })

        Convey("properly decodes paths", func() {
            p, err := plugins.NewDiskUsage(plugins.Params{"Paths": []string{"foo", "bar"}})
            So(p, ShouldNotBeNil)
            So(err, ShouldBeNil)
        })
    })
}
