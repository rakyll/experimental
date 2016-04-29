package gpio

type Device struct{}

func Open() (*Device, error) {
	panic("not yet implemented")
}

func (d *Device) Read(pin int) (int, error) {
	panic("not yet implemented")
}

func (d *Device) Write(pin int, v int) error {
	panic("not yet implemented")
}

func (d *Device) SetActiveType(pin int, typ int) error {
	panic("not yet implemented")
}

func (d *Device) SetDirection(pin int, direction int) error {
	panic("not yet implemented")
}

func (d *Device) SetEdgeTriggerType(pin int, typ int) error {
	panic("not yet implemented")
}

// TODO(jbd): Allow polling.

func (d *Device) Close() error {
	panic("not yet implemented")
}
