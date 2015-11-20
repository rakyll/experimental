package spi

type Device struct {
	id string
}

func (d *Device) Read(buf []byte) (int, error) {
	panic("not implemented")
}

func (d *Device) Write(buf []byte) error {
	panic("not implemented")
}

func (d *Device) SetActive(id int) {
	panic("not implemented")
}

func (d *Device) Close() error {
	panic("not implemented")
}

func (d *Device) SetSpeed(speedHz int) error {
	panic("not implemented")
}

func (d *Device) SetDelay(delay int) error {
	panic("not implemented")
}

func Open(device string, mode int, speedHz int, delay int) *Device {
	panic("not implemented")
}
