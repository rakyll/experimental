package gpio

import "github.com/rakyll/experimental/gpio/driver"

type Device struct {
	conn driver.Conn
}

func Open(d driver.Opener) (*Device, error) {
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
	return d.conn.SetValue(pin, v)
}

func (d *Device) SetActiveType(pin int, typ int) error {
	return d.conn.SetActiveType(pin, typ)
}

func (d *Device) SetDirection(pin int, direction int) error {
	return d.conn.SetDirection(pin, direction)
}

func (d *Device) SetEdgeTriggerType(pin int, typ int) error {
	return d.conn.SetEdgeTriggerType(pin, typ)
}

// TODO(jbd): Allow polling.

func (d *Device) Close() error {
	return d.conn.Close()
}
