package driver

type Opener interface {
	Open() (Conn, error)
}

type Conn interface {
	Val(pin int) (int, error)
	SetVal(pin int, v int) error
	SetActiveType(pin int, t string) error
	SetDirection(pin int, t string) error
	SetEdgeTriggerType(pin int, t string) error
	Close() error
}
