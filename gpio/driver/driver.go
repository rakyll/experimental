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
	Val(pin int) (int, error)
	SetVal(pin int, v int) error
	Configure(pin int, opt PinOptions) error
	Close() error
}
