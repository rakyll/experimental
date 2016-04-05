package i2c

import (
	"github.com/rakyll/experimental/i2c/driver"
)

type DevFS struct{}

func (d *DevFS) Open(addr, bus int) (driver.Conn, error) {
	panic("not yet implemented")
}

type devFSConn struct{}

func (c *devFSConn) Read(buf []byte) (int, error) {
	panic("not yet implemented")
}

func (c *devFSConn) Write(buf []byte) error {
	panic("not yet implemented")
}

func (c *devFSConn) Close() error {
	panic("not yet implemented")
}
