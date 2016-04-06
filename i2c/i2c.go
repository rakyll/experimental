package i2c

import (
	"github.com/rakyll/experimental/i2c/driver"
)

type Device struct {
	conn driver.Conn
}

func (d *Device) Read(buf []byte) (int, error) {
	return d.conn.Read(buf)
}

func (d *Device) Write(buf []byte) (int, error) {
	return d.conn.Write(buf)
}

func (d *Device) Close() error {
	return d.conn.Close()
}

// Open opens an I2C device with the given I2C address and adapter number.
func Open(o driver.Opener, addr, adapter int) (*Device, error) {
	conn, err := o.Open(addr, bus)
	if err != nil {
		return nil, err
	}
	return &Device{conn: conn}, nil
}
