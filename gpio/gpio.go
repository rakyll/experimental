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

type EdgeTriggerType string

const (
	None    = EdgeTriggerType("none")
	Rising  = EdgeTriggerType("rising")
	Falling = EdgeTriggerType("falling")
	Both    = EdgeTriggerType("both")
)

type ActiveType string

const (
	ActiveLow  = ActiveType("active_low")
	ActiveHigh = ActiveType("active_high")
)

func Open(d driver.Opener) (*Device, error) {
	// TODO(jbd): Open pin rather than GPIO device? It would help
	// some driver implementations such as sysfs.
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
	return d.conn.Configure(pin, string(dir), "", "")
}

func (d *Device) Configure(pin int, dir Direction, a ActiveType, e EdgeTriggerType) error {
	return d.conn.Configure(pin, string(dir), string(a), string(e))
}

// TODO(jbd): Allow polling the value of a pin.

func (d *Device) Close() error {
	return d.conn.Close()
}
