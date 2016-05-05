package gpio

import "github.com/rakyll/experimental/gpio/driver"

type Devfs struct {
	// BaseAddr represents the base memory address of the BCM2835 controller.
	// Given the zero value, Devfs will try to find the correct value by looking
	// at /proc/device-tree/soc/ranges.
	BaseAddr uint64
}

func (d *Devfs) Open() (driver.Conn, error) {
	panic("not implemented")
}

type devfsConn struct{}

func (d *devfsConn) Value(pin int) (int, error) {
	panic("not implemented")
}

func (d *devfsConn) SetValue(pin int, v int) error {
	panic("not implemented")
}

func (d *devfsConn) SetDirection(pin int, dir string) error {
	panic("not implemented")
}

func (d *devfsConn) SetActiveLow(pin int, low bool) error {
	panic("not implemented")
}

func (d *devfsConn) Close() error {
	panic("not implemented")
}
