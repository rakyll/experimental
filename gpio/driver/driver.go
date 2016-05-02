package driver

type PinOptions struct {
	ActiveType      string
	Direction       string
	EdgeTriggerType string
}

type Opener interface {
	Open() (Conn, error)
}

type Conn interface {
	Configure(pin int, opt PinOptions) error
	Value(pin int) (int, error)
	SetValue(pin int, v int) error
	Close() error
}
