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
	"fmt"
	"math"

	"github.com/rakyll/experimental/spi"
)

func main() {
	d, err := spi.Open("/dev/spidev0.1")
	if err != nil {
		panic(err)
	}
	defer d.Close()

	fmt.Println("opened spidev", d)

	if err := d.SetMode(spi.Mode3); err != nil {
		panic(err)
	}
	if err := d.SetSpeed(1000000); err != nil {
		panic(err)
	}
	if err := d.SetBitsPerWord(8); err != nil {
		panic(err)
	}
	drawRainbow(d, 0.7)
}

// n is the number of LEDs in the strip.
const n = 50

func drawRainbow(dev *spi.Device, freq float64) {
	payload := make([]byte, n*4+n/2+1)
	payload[0] = 0
	payload[1] = 0
	payload[2] = 0
	payload[3] = 0

	for i := 0; i < n; i++ {
		v := float64(i)
		payload[(i+1)*4] = 0xff                                 // alpha
		payload[(i+1)*4+1] = byte(math.Sin(freq*v+0)*127 + 128) // red
		payload[(i+1)*4+2] = byte(math.Sin(freq*v+2)*127 + 128) // green
		payload[(i+1)*4+3] = byte(math.Sin(freq*v+4)*127 + 128) // blue
	}
	if err := dev.Do(payload, 0); err != nil {
		panic(err)
	}
}
