package audio

type Device struct{}

func OpenDevice(name string) (*Device, error) {
	panic("not yet")
}

func (d *Device) OpenChannel() *Channel {
	panic("not yet implemented")
}

func (d *Device) Close() error {
	panic("not yet implemented")
}

type Channel struct {
	d *Device
}

func (c *Channel) Write(b []byte) (int, error) {
	panic("not yet implemented")
}

func (c *Channel) Close() error {
	panic("not yet implemented")
}