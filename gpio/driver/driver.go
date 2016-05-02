package driver

type Opener interface {
	Open() (Conn, error)
}

type Conn interface {
	Value(pin int) (int, error)
	SetValue(pin int, v int) error
	Configure(pin int, direction, active, edgeTrigger string) error
	Close() error
}
