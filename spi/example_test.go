// Copyright 2014 Google Inc. All Rights Reserved.
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

package spi_test

import "github.com/rakyll/experimental/spi"

func Example() {
	dev, err := spi.Open("/dev/spidev0.1")
	if err != nil {
		panic(err)
	}
	defer dev.Close()

	if err := dev.SetMode(spi.Mode3); err != nil {
		panic(err)
	}
	if err := dev.SetSpeed(500000); err != nil {
		panic(err)
	}
	if err := dev.SetBitsPerWord(8); err != nil {
		panic(err)
	}
	if err := dev.Do([]byte{
		0, 0, 0, 0,
		0xff, 200, 0, 200,
		0xff, 200, 0, 200,
		0xe0, 200, 0, 200,
		0xff, 200, 0, 200,
		0xff, 8, 50, 0,
		0xff, 200, 0, 0,
		0xff, 0, 0, 0,
		0xff, 200, 0, 200,
		0xff, 0xff, 0xff, 0xff,
	}, 0); err != nil {
		panic(err)
	}
}
