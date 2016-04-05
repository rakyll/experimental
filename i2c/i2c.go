package i2c

import (
	"github.com/rakyll/experimental/i2c/driver"
)

type Device struct {
	conn driver.Conn
}

func Open(o driver.Opener, addr, bus int) (*Device, error) {
	panic("not implemented yet")
}
