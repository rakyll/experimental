package driver

type Opener interface {
	Open() (Conn, error)
}

type Conn interface {
	Value(pin int) (int, error)
	SetValue(pin int, v int) error
	SetActiveType(pin int, t string) error
	SetDirection(pin int, t string) error
	SetEdgeTriggerType(pin int, t string) error
	Close() error
}
