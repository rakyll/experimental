package gpio

import "github.com/rakyll/experimental/gpio/driver"

type Device struct {
	conn driver.Conn
}

type Direction string

const (
	In  = Direction("in")
	Out = Direction("out")
	// TODO(jbd): Out but initially high or initially low?
)

// TODO(jbd): How to support analog pins?

// TODO(jbd): Allow users to configure edge trigger type.

// Open opens a connection to the GPIO device with the given driver.
// If a nil driver is provided, it uses the Devfs implementation with
// the default settings.
func Open(d driver.Opener) (*Device, error) {
	// TODO(jbd): Open pin rather than GPIO device? It would help
	// some driver implementations such as sysfs.
	if d == nil {
		d = &Devfs{}
	}
	conn, err := d.Open()
	if err != nil {
		return nil, err
	}
	return &Device{conn: conn}, nil
}

func (d *Device) Value(pin int) (int, error) {
	return d.conn.Value(pin)
}

func (d *Device) SetValue(pin int, v int) error {
	return d.conn.SetValue(pin, int(v))
}

func (d *Device) SetDirection(pin int, dir Direction) error {
	return d.conn.SetDirection(pin, string(dir))
}

func (d *Device) SetActiveLow(pin int, low bool) error {
	return d.conn.SetActiveLow(pin, low)
}

func (d *Device) Close() error {
	return d.conn.Close()
}
