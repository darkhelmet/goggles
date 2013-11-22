package config_test

import (
    "github.com/darkhelmet/goggles/config"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestLoadFile(t *testing.T) {
    Convey("Loading", t, func() {
        Convey("a basic file", func() {
            c, err := config.LoadFile("test_data/basic.json")
            So(c, ShouldNotBeNil)
            So(err, ShouldBeNil)

            So(c.Hostname, ShouldEqual, "atlas")

            So(c.InfluxDB.Host, ShouldEqual, "localhost")
            So(c.InfluxDB.Port, ShouldEqual, 8086)
            So(c.InfluxDB.Database, ShouldEqual, "goggles")
            So(c.InfluxDB.Username, ShouldEqual, "brucewayne")
            So(c.InfluxDB.Password, ShouldEqual, "batman")
        })

        Convey("without InfluxDB info", func() {
            c, err := config.LoadFile("test_data/name.json")
            So(c, ShouldBeNil)
            So(err, ShouldNotBeNil)
        })

        Convey("with InfluxDB defaults", func() {
            c, err := config.LoadFile("test_data/influxdb_defaults.json")
            So(c, ShouldNotBeNil)
            So(err, ShouldBeNil)

            So(c.InfluxDB.Port, ShouldEqual, 8086)
            So(c.InfluxDB.Database, ShouldEqual, "goggles")
        })
    })
}
