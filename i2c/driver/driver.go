// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package driver contains interfaces to be implemented by various I2C implementations.
package driver

// Opener is an interface to be implemented by the I2C driver to open
// a connection an I2C device with the specified addr and bus number.
type Opener interface {
	Open(addr uint8, bus int) (Conn, error)
}

// Conn represents an active connection to an I2C device.
type Conn interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Close() error
}
