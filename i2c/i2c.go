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

func Open(o driver.Opener, addr uint8, bus int) (*Device, error) {
	if o == nil {
		o = &DevFS{}
	}
	conn, err := o.Open(addr, bus)
	if err != nil {
		return nil, err
	}
	return &Device{conn: conn}, nil
}
