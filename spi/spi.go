package spi

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	IOC_MAGIC = 107

	IOC_NRBITS   = 8
	IOC_TYPEBITS = 8
	IOC_SIZEBITS = 13 /* Actually 14, see below. */
	IOC_DIRBITS  = 3

	IOC_NRMASK   = (1 << IOC_NRBITS) - 1
	IOC_TYPEMASK = (1 << IOC_TYPEBITS) - 1
	IOC_SIZEMASK = (1 << IOC_SIZEBITS) - 1
	IOC_DIRMASK  = (1 << IOC_DIRBITS) - 1

	IOC_NRSHIFT   = 0
	IOC_TYPESHIFT = IOC_NRSHIFT + IOC_NRBITS
	IOC_SIZESHIFT = IOC_TYPESHIFT + IOC_TYPEBITS
	IOC_DIRSHIFT  = IOC_SIZESHIFT + IOC_SIZEBITS

	IOC_NONE  = 0
	IOC_READ  = 2
	IOC_WRITE = 4
)

func ioc(dir, typ, nr, size uintptr) uintptr {
	return (dir << IOC_DIRSHIFT) | (typ << IOC_TYPESHIFT) | (nr << IOC_NRSHIFT) | (size << IOC_SIZESHIFT)
}

func wrIOC() uintptr {
	return ioc(IOC_READ|IOC_WRITE, IOC_MAGIC, 1, 1)
}

func speedHzIOC() uintptr {
	panic("not implemented")
}

type Device struct {
	f           *os.File
	mode        int
	delay       int
	speedHz     int
	bitsPerWord int
}

type payload struct {
	tx          uint64
	rx          uint64
	length      uint32
	speedHz     uint32
	delay       uint16
	bitsPerWord uint8
}

func (d *Device) SetMode(mode int) error {
	if err := d.ioctl(wrIOC(), uintptr(unsafe.Pointer(&mode))); err != nil {
		return err
	}
	d.mode = mode
	return nil
}

func (d *Device) SetSpeed(speedHz int) error {
	if err := d.ioctl(speedHzIOC(), uintptr(unsafe.Pointer(&speedHz))); err != nil {
		return err
	}
	d.speedHz = speedHz
	return nil
}

func (d *Device) SetBitsPerWord(bits int) error {
	panic("not implemented")
}

func (d *Device) Write(data []byte) error {
	panic("not implemented")
}

func (d *Device) Close() error {
	return d.f.Close()
}

func (d *Device) ioctl(a1, a2 uintptr) error {
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL, d.f.Fd(), a1, a2,
	)
	if errno != 0 {
		return syscall.Errno(errno)
	}
	return nil
}

func Open(device string, mode int, speed int, bitsPerWord int) (*Device, error) {
	f, err := os.Open(device)
	if err != nil {
		return nil, err
	}
	d := &Device{f: f}
	if err := d.SetMode(mode); err != nil {
		d.Close()
		return nil, err
	}
	if err := d.SetSpeed(speed); err != nil {
		d.Close()
		return nil, err
	}
	if err := d.SetBitsPerWord(bitsPerWord); err != nil {
		d.Close()
		return nil, err
	}
	// TODO(jbd): close if error
	return d, nil
}
