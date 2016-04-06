package i2c

import (
	"fmt"
	"os"
	"syscall"

	"github.com/rakyll/experimental/i2c/driver"
)

type DevFS struct{}

func (d *DevFS) Open(addr, bus int) (driver.Conn, error) {
	f, err := os.OpenFile(fmt.Sprintf("/dev/i2c-%d", bus), os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	conn := &devFSConn{f: f}
	if err := conn.ioctl(0x0703, uintptr(addr)); err != nil {
		conn.Close()
		return nil, err
	}
	return &devFSConn{f: f}, nil
}

type devFSConn struct {
	f *os.File
}

func (c *devFSConn) Read(buf []byte) (int, error) {
	return c.f.Read(buf)
}

func (c *devFSConn) Write(buf []byte) (int, error) {
	return c.f.Write(buf)
}

func (c *devFSConn) Close() error {
	return c.f.Close()
}

func (c *devFSConn) ioctl(arg1, arg2 uintptr) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, c.f.Fd(), arg1, arg2)
	if errno != 0 {
		return syscall.Errno(errno)
	}
	return nil
}
