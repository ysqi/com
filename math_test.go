// Copyright 2015 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package com

import (
	"math"
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Pow(t *testing.T) {
	Convey("Power int", t, func() {
		for x := 0; x < 10; x++ {
			for y := 0; y < 8; y++ {
				result := PowInt(x, y)
				result_float := math.Pow(float64(x), float64(y))
				So(result, ShouldEqual, int(result_float))
			}
		}
	})
}

func BenchmarkPow(b *testing.B) {
	x := rand.Intn(100)
	y := rand.Intn(6)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		PowInt(x, y)
	}
}

func Test_Round(t *testing.T) {
	Convey("Round float64", t, func() {
		testCaess := map[float64]float64{
			+0.00: 0.00,
			+0.50: 1.00,
			-0.50: 0.00,
			+1.23: 1.00,
			-1.23: -1.00,
			+1.25: 1,
			-1.25: -1.00,
			+2.00: 2,
			-2.00: -2.00,
			25.00: 25,
			25.55: 26.00,

			123.555555:          124,
			123.333333333333333: 123,
		}
		for v, want := range testCaess {
			got := Round(v)
			So(got, ShouldEqual, want)
		}
	})
}

func Test_RoundFloat(t *testing.T) {
	Convey("RoundFloat float64", t, func() {
		testCaess := map[float64]float64{
			+0.00:  0.00,
			+0.50:  0.50,
			-0.50:  -0.50,
			+1.23:  1.23,
			-1.23:  -1.23,
			+1.25:  1.25,
			-1.25:  -1.25,
			+2.00:  2,
			-2.00:  -2.00,
			25.00:  25.00,
			25.555: 25.56,

			123.555555:          123.56,
			123.333333333333333: 123.33,
			-1.2456:             -1.25,
		}
		for v, want := range testCaess {
			got := RoundFloat(v, 2)
			So(got, ShouldEqual, want)
		}
	})
}
