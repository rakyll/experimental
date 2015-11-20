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

func Open(device string, mod int, clock int) *Device {
	panic("not implemented")
}
