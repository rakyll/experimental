package spi

import (
	"os"
	"syscall"
	"time"
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

	nrbits   = 8
	typebits = 8
	sizebits = 13 /* Actually 14, see below. */
	dirbits  = 3

	nrshift   = 0
	typeshift = nrshift + nrbits
	sizeshift = typeshift + typebits
	dirshift  = sizeshift + sizebits

	none  = 0
	read  = 2
	write = 4
)

func ioc(dir, typ, nr, size uintptr) uintptr {
	return (dir << dirshift) | (typ << typeshift) | (nr << nrshift) | (size << sizeshift)
}

func wrIOC() uintptr {
	return ioc(write, magic, 1, 1)
}

func bitsPerWordIOC() uintptr {
	return ioc(write, magic, 3, 1)
}

func speedHzIOC() uintptr {
	return ioc(write, magic, 4, 4)
}

func msgIOC(n uint32) uintptr {
	return uintptr(0x40006B00 + (n * 0x200000))
}

type Device struct {
	f           *os.File
	mode        uint8
	speedHz     uint32
	bitsPerWord uint8
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
	d.mode = m
	return nil
}

func (d *Device) SetSpeed(speedHz int) error {
	s := uint32(speedHz)
	if err := d.ioctl(speedHzIOC(), uintptr(unsafe.Pointer(&s))); err != nil {
		return err
	}
	d.speedHz = s
	return nil
}

func (d *Device) SetBitsPerWord(bits int) error {
	b := uint8(bits)
	if err := d.ioctl(bitsPerWordIOC(), uintptr(unsafe.Pointer(&b))); err != nil {
		return err
	}
	d.bitsPerWord = b
	return nil
}

func (d *Device) Do(buf []byte, delay time.Duration) error {
	p := payload{
		tx:          uint64(uintptr(unsafe.Pointer(&buf[0]))),
		rx:          uint64(uintptr(unsafe.Pointer(&buf[0]))),
		length:      uint32(len(buf)),
		speedHz:     d.speedHz,
		delay:       uint16(delay),
		bitsPerWord: d.bitsPerWord,
	}
	// TODO(jbd): read from the buffer.
	return d.ioctl(msgIOC(1), uintptr(unsafe.Pointer(&p)))
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

func Open(name string) (*Device, error) {
	f, err := os.OpenFile(name, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	return &Device{f: f}, nil
}
