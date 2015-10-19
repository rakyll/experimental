// +build darwin
// +build arm arm64

package player

import "unsafe"

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreMotion
#import <stdlib.h>
*/
import "C"

type Player struct {
	p unsafe.Pointer
}

func NewPlayer(url string) *Player {
	p := newPlayer(nil)
	return &Player{p: p}
}

func (p *Player) Play() {
	play(p.p)
}

func (p *Player) Pause() {
	pause(p.p)
}
