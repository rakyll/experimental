package spi

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	Mode0 = 0x00
	Mode1 = 0x04
	Mode2 = 0x08
	Mode3 = 0x0C
)

const (
	magic = 107

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
	return ioc(IOC_WRITE, magic, 1, 1)
}

func bitsPerWordIOC() uintptr {
	return ioc(IOC_WRITE, magic, 3, 1)
}

func speedHzIOC() uintptr {
	return ioc(IOC_WRITE, magic, 4, 4)
}

func msgIOC(n uint32) uintptr {
	return uintptr(0x40006B00 + (n * 0x200000))
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
	m := uint8(mode)
	if err := d.ioctl(wrIOC(), uintptr(unsafe.Pointer(&m))); err != nil {
		return err
	}
	d.mode = mode
	return nil
}

func (d *Device) SetSpeed(speedHz int) error {
	s := uint32(speedHz)
	if err := d.ioctl(speedHzIOC(), uintptr(unsafe.Pointer(&s))); err != nil {
		return err
	}
	d.speedHz = speedHz
	return nil
}

func (d *Device) SetBitsPerWord(bits int) error {
	b := uint8(bits)
	if err := d.ioctl(bitsPerWordIOC(), uintptr(unsafe.Pointer(&b))); err != nil {
		return err
	}
	d.bitsPerWord = bits
	return nil
}

func (d *Device) Do(tx, rx []byte) error {
	p := payload{
		tx:          uint64(uintptr(unsafe.Pointer(&tx))),
		rx:          uint64(uintptr(unsafe.Pointer(&rx))),
		length:      uint32(len(tx)),
		speedHz:     uint32(d.speedHz),
		delay:       uint16(d.delay),
		bitsPerWord: uint8(d.bitsPerWord),
	}
	if err := d.ioctl(msgIOC(), uintptr(unsafe.Pointer(&p))); err != nil {
		return err
	}
	// TODO(jbd): Read into rx.
	return nil
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

func Open(device string) (*Device, error) {
	f, err := os.OpenFile(device, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	return &Device{f: f}, nil
}
