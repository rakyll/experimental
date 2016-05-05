package gpio

type Devfs struct {
	// BaseAddr represents the base memory address of the BCM2835 controller.
	// Given the zero value, Devfs will try to find the correct value by looking
	// at /proc/device-tree/soc/ranges.
	BaseAddr uint64
}

func (d *Devfs) Value(pin int) (int, error) {
	panic("not implemented")
}

func (d *Devfs) SetValue(pin int, v int) error {
	panic("not implemented")
}

func (d *Devfs) Configure(pin int, direction, active, edgeTrigger string) error {
	panic("not implemented")
}

func (d *Devfs) Close() error {
	panic("not implemented")
}
