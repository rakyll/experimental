package driver

type Opener interface {
	Open() (Conn, error)
}

type Conn interface {
	Value(pin int) (int, error)
	SetValue(pin int, v int) error
	SetActiveType(pin int, typ int) error
	SetDirection(pin int, direction int) error
	SetEdgeTriggerType(pin int, typ int) error
	Close() error
}
