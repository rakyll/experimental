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

// Decoder decodes the source into a Clip.
func Decode(src io.ReadSeeker) (Clip, error) {
	panic("not implemented")
}

// NewBufferClip converts a buffer to a Clip.
func NewBufferClip(buf []byte, info FrameInfo) Clip {
	panic("not implemented")
}

// NewRemoteClip convert the HTTP live streaming media
// source into a Clip.
func NewRemoteClip(url string) (Clip, error) {
	panic("not implemented")
}

// FrameInfo represents the frame-level information.
type FrameInfo struct {
	// Channels represent the number of audio channels (e.g. 1 for mono, 2 for stereo).
	Channels int

	// Bit depth is the number of bits used to represent a single sample.
	BitDepth int

	// Sample rate is the number of samples to be played each second.
	SampleRate int64
}

// Clip represents a linear PCM formatted audio io.ReadSeeker.
// Clip can seek and read from a section and allow users to
// consume a small section of the underlying audio data.
// FrameInfo returns the basic frame-level information about the clip audio.
// Size returns the total number of bytes from the underlying audio
// data. If size is not known, -1 is returned.
type Clip interface {
	io.ReadSeeker
	FrameInfo() FrameInfo
}

// State indicates the current playing state of the player.
type State int

type Player struct{}

// NewPlayer returns a new Player.
// (*Player).Play plays the provided clips in the given order.
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
