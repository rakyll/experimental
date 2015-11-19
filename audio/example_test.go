package audio_test

import (
	"io"

	"github.com/rakyll/experimental/audio"
)

func DecodeClip(r io.ReadSeeker) audio.Clip {
	return &decodedClip{in: r}
}

func Example() {
	var rs io.ReadSeeker
	clip := DecodeClip(rs)

	p, err := audio.NewPlayer(clip)
	if err != nil {
		panic(err)
	}
	p.Play()

}

type decodedClip struct {
	in io.Reader

	current int64  // current is the current frame offset. It is the number of all consumed frames.
	buf     []byte // TODO(jbd): Buffer is a small buffer to be used for prefetching and caching.
}

func (s *decodedClip) Seek(offset int64, whence int) (int64, error) {
	panic("not implemented")
}

func (s *decodedClip) Read(buf []byte) (n int, err error) {
	// if buffer contains at least len(buf) bytes, return them.

	// if buffer is not filled with enough data, read more from
	// s.in.

	// if s.in is EOFed, return EOF.
	panic("not yet implemented")
}

func (s *decodedClip) Info() audio.ClipInfo {
	// TODO(jbd): Determined from the header from s.in.
	return audio.ClipInfo{2, 16, 44000}
}
