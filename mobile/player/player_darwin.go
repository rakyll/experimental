package player

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework AVFoundation

#import <AVFoundation/AVPlayer.h>
#import <CoreFoundation/CoreFoundation.h>

AVPlayer* newPlayer(char* url);

*/
import "C"

type Player struct {
	p *C.struct_AVPlayer
}

func NewPlayer(url string) *Player {
	p := C.newPlayer(nil)
	return &Player{p: p}
}
