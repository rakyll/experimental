package gpio

import "github.com/rakyll/experimental/gpio/driver"

type Device struct {
	conn driver.Conn
}

type Value int

const (
	Low  = Value(0)
	High = Value(1)
)

type Direction string

const (
	In               = Direction("in")
	OutInitiallyLow  = Direction("out_low")
	OutInitiallyHigh = Direction("out_high")
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
	conn, err := d.Open()
	if err != nil {
		return nil, err
	}
	return &Device{conn: conn}, nil
}

func (d *Device) Val(pin int) (Value, error) {
	v, err := d.conn.Val(pin)
	if err != nil {
		return Value(0), err
	}
	return Value(v), nil
}

func (d *Device) SetVal(pin int, v Value) error {
	return d.conn.SetVal(pin, int(v))
}

func (d *Device) SetActiveType(pin int, t ActiveType) error {
	return d.conn.Configure(pin, driver.PinOptions{
		ActiveType: string(t),
	})
}

func (d *Device) SetDirection(pin int, dir Direction) error {
	return d.conn.Configure(pin, driver.PinOptions{
		Direction: string(dir),
	})
}

func (d *Device) SetEdgeTriggerType(pin int, t EdgeTriggerType) error {
	return d.conn.Configure(pin, driver.PinOptions{
		EdgeTriggerType: string(t),
	})
}

// TODO(jbd): Allow polling.

func (d *Device) Close() error {
	return d.conn.Close()
}
