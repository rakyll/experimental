package driver

type Opener interface {
	Open(addr uint8, bus int) (Conn, error)
}

type Conn interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Close() error
}
