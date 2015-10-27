package audio

type Device struct{}

func (d *Device) OpenChannel() *Channel {
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
