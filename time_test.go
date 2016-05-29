package com

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetFirstWeekDate(t *testing.T) {
	taseCase := []struct {
		Test string
		Want string
	}{
		{Test: "2016-05-29", Want: "2016-05-23"},
		{Test: "2016-05-28", Want: "2016-05-23"},
		{Test: "2016-05-30", Want: "2016-05-30"},
		{Test: "2016-05-23", Want: "2016-05-23"},
		{Test: "2016-05-22", Want: "2016-05-16"},
		{Test: "2016-05-21", Want: "2016-05-16"},
		{Test: "2016-05-20", Want: "2016-05-16"},
		{Test: "2016-05-19", Want: "2016-05-16"},
		{Test: "2016-05-17", Want: "2016-05-16"},
		{Test: "2016-05-16", Want: "2016-05-16"},
		{Test: "2016-05-15", Want: "2016-05-09"},
		{Test: "000-00-00", Want: "0000-000-00"},
	}
	Convey("Test GetFirstWeekDate", t, func() {
		for _, v := range taseCase {
			test, _ := DateParse(v.Test, "Y-m-d")
			want, _ := DateParse(v.Want, "Y-m-d")
			got := FirstDayOfWeek(test)
			So(got.Unix(), ShouldEqual, want.Unix())
		}
	})
}
