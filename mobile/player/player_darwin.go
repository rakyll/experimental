package player

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework AVFoundation

#import <stdlib.h>

#import <AVFoundation/AVPlayer.h>
#import <CoreFoundation/CoreFoundation.h>

AVPlayer* newPlayer(char* url);
void playPlayer(AVPlayer*);
void pausePlayer(AVPlayer*);

*/
import "C"

import "unsafe"

type Player struct {
	p *C.struct_AVPlayer
}

func NewPlayer(url string) *Player {
	cu := C.CString(url)
	defer C.free(unsafe.Pointer(cu))

	p := C.newPlayer(cu)
	return &Player{p: p}
}

func (p *Player) Play() {
	C.playPlayer(p.p)
}

func (p *Player) Pause() {
	C.pausePlayer(p.p)
}
