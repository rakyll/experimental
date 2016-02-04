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
	if err := d.SetSpeed(100000); err != nil {
		panic(err)
	}
	if err := d.SetBitsPerWord(8); err != nil {
		panic(err)
	}
	if err := d.Do([]byte{
		0, 0, 0, 0,
		0xff, 0xee, 0xee, 0xff,
		0xff, 0xee, 0xee, 0xff,
		0xff, 0xee, 0xee, 0xff,
		0xff, 0xee, 0xee, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff,
	}, 0); err != nil {
		panic(err)
	}
}
