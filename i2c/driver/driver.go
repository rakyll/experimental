package driver

type Opener interface {
	Open(addr, bus int) (Conn, error)
}

type Conn interface {
	Read(buf []byte) (int, error)
	Write(buf []byte) error
	Close() error
}
