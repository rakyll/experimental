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
	"bytes"
	"errors"
	"fmt"
	"io"
	"sync"
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

func (*Decoder) Stream() Stream {}

func NewMP3Decoder(src Stream) *Decoder {
	panic("not yet")
}

func NewDecoder(src Stream) *Decoder {
	panic("not yet")
}

func a() {
	var s Stream
	d := NewMP3Decoder(s)
	d2 := NewDecoder(d.Decoded())
	d3 := NewDecoder(d2.Decoded())
	d3.Decoded()
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

// Player is a basic audio player that plays PCM data.
// Operations on a nil *Player are no-op, a nil *Player can
// be used for testing purposes.
type Player struct{}

// NewPlayer returns a new Player.
// It initializes the underlying audio devices and the related resources.
// If zero values are provided for format and sample rate values, the player
// determines them from the source's WAV header.
// An error is returned if the format and sample rate can't be determined.
//
// The audio package is only designed for small audio sources.
func NewPlayer(src ...Stream) (*Player, error) {
	panic("not yet implemented")
}

// Play buffers the source audio to the audio device and starts
// to play the source.
// If the player paused or stopped, it reuses the previously buffered
// resources to keep playing from the time it has paused or stopped.
func (p *Player) Play() error {
	panic("not yet implemented")
}

// Pause pauses the player.
func (p *Player) Pause() error {
	panic("not yet implemented")
}

// Stop stops the player.
func (p *Player) Stop() error {
	panic("not yet implemented")
}

// Seek moves the play head to the given offset relative to the start of the source.
func (p *Player) Seek(offset time.Duration) error {
	panic("not yet implemented")
}

// Current returns the current playback position of the audio that is being played.
func (p *Player) Current() time.Duration {
	panic("not yet implemented")
}

// Total returns the total duration of the audio source.
func (p *Player) Total() time.Duration {
	panic("not yet implemented")
}

// Volume returns the current player volume. The range of the volume is [0, 1].
func (p *Player) Volume() float64 {
	panic("not yet implemented")
}

// SetVolume sets the volume of the player. The range of the volume is [0, 1].
func (p *Player) SetVolume(vol float64) {
	panic("not yet implemented")
}

// State returns the player's current state.
func (p *Player) State() State {
	panic("not yet implemented")
}

// Close closes the device and frees the underlying resources
// used by the player.
// It should be called as soon as the player is not in-use anymore.
func (p *Player) Close() error {
	panic("not yet implemented")
}
