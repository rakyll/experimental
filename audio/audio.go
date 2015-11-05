// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package audio provides a basic audio player.
//
// In order to use this package on Linux desktop distros,
// you will need OpenAL library as an external dependency.
// On Ubuntu 14.04 'Trusty', you may have to install this library
// by running the command below.
//
//    sudo apt-get install libopenal-dev
//
// When compiled for Android, this package uses OpenAL Soft as a backend.
// Please add its license file to the open source notices of your
// application.
// OpenAL Soft's license file could be found at
// http://repo.or.cz/w/openal-soft.git/blob/HEAD:/COPYING.
package audio

import (
	"time"

	"golang.org/x/mobile/exp/audio/al"
)

func NewRemoteStream(url string) Stream {
	panic("not implemented")
}

type Stream interface {
	Read(n int) ([]byte, int, error)
	Frame() (channels int, bitDepth int64, samplesPerSecond int64)
}

type Decoder struct{}

func (*Decoder) Stream() Stream {
	panic("not implemented")
}

func NewMP3Decoder(src Stream) *Decoder {
	panic("not yet")
}

func NewDecoder(src Stream) *Decoder {
	panic("not yet")
}

func a() {
	var s Stream
	d := NewMP3Decoder(s)
	d2 := NewDecoder(d.Stream())
	d3 := NewDecoder(d2.Stream())
	d3.Stream()
	// decoding starts when you start to read from d3 decoded stream.
}

// State indicates the current playing state of the player.
type State int

const (
	Unknown State = iota
	Initial
	Playing
	Paused
	Stopped
)

func (s State) String() string { return stateStrings[s] }

var stateStrings = [...]string{
	Unknown: "unknown",
	Initial: "initial",
	Playing: "playing",
	Paused:  "paused",
	Stopped: "stopped",
}

var codeToState = map[int32]State{
	0:          Unknown,
	al.Initial: Initial,
	al.Playing: Playing,
	al.Paused:  Paused,
	al.Stopped: Stopped,
}

type Player struct{}

func NewPlayer(src ...Stream) (*Player, error) {
	panic("not yet implemented")
}

func (p *Player) Play() error {
	panic("not yet implemented")
}

func (p *Player) Pause() error {
	panic("not yet implemented")
}

func (p *Player) Stop() error {
	panic("not yet implemented")
}

func (p *Player) Seek(offset time.Duration) error {
	panic("not yet implemented")
}

func (p *Player) Current() time.Duration {
	panic("not yet implemented")
}

func (p *Player) Total() time.Duration {
	panic("not yet implemented")
}

func (p *Player) Volume() float64 {
	panic("not yet implemented")
}

func (p *Player) SetVolume(vol float64) {
	panic("not yet implemented")
}

func (p *Player) State() State {
	panic("not yet implemented")
}

func (p *Player) Close() error {
	panic("not yet implemented")
}
