package driver

type Opener interface {
	Open() (Conn, error)
}

type Conn interface {
	Value(pin int) (int, error)
	SetValue(pin int, v int) error
	SetDirection(pin int, dir string) error
	SetActiveLow(pin int, low bool) error
	Close() error
}
