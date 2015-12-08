// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"math"
	"time"

	"github.com/rakyll/experimental/spi"
)

func main() {
	d, err := spi.Open("/dev/spidev0.1")
	if err != nil {
		panic(err)
	}
	defer d.Close()

	if err := d.SetMode(spi.Mode3); err != nil {
		panic(err)
	}
	if err := d.SetSpeed(10000); err != nil {
		panic(err)
	}
	if err := d.SetBitsPerWord(8); err != nil {
		panic(err)
	}
	for {
		loop(d)
	}
}

// n is the number of LEDs in the strip.
const n = 19

var (
	alpha byte = 0xe0
	delta byte = 0x1
)

func loop(dev *spi.Device) {
	colors := make([]byte, 4*n)
	freq := 0.7
	for i := float64(0); i < n; i++ {
		red := byte(math.Sin(freq*i+0)*127 + 128)
		green := byte(math.Sin(freq*i+2)*127 + 128)
		blue := byte(math.Sin(freq*i+4)*127 + 128)
		colors = append(colors, alpha, red, green, blue)
	}

	alpha += delta
	if alpha == 0xe0 || alpha == 0xff {
		delta = -delta
	}

	payload := []byte{0, 0, 0, 0}
	payload = append(payload, colors...)
	payload = append(payload, 0xff, 0xff, 0xff, 0xff)

	if err := dev.Do(payload, 0); err != nil {
		panic(err)
	}
	time.Sleep(100 * time.Millisecond)
}
