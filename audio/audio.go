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
	"io"
	"time"
)

func NewBufferClip(buf []byte) Clip {
	panic("not implemented")
}

func NewRemoteStream(url string) (Clip, error) {
	panic("not implemented")
}

type ClipInfo struct {
	Channels   int
	BitDepth   int64
	SampleRate int64
}

type Clip interface {
	io.ReadSeeker
	Info() ClipInfo
}

// State indicates the current playing state of the player.
type State int

type Player struct{}

func NewPlayer(c ...Clip) (*Player, error) {
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
