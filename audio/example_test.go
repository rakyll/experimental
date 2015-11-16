package audio_test

import "github.com/rakyll/experimental/audio"

func Example() {
	var r io.Reader
	s, err := Decode(r)
	if err != nil {
		panic(err)
	}

	// decoding begins when d3.Ouput needs to be consumed.
	p, err := audio.NewPlayer(s)
	if err != nil {
		panic(err)
	}
	p.Play()

}

func Decode(r io.Reader) Stream {
	return decodedStream{in: r}
}

type decodedStream struct {
	in io.Reader

	current int64  // current is the current frame offset. It is the number of all consumed frames.
	buf     []byte // TODO(jbd): Buffer is a small buffer to be used for prefetching and caching.
}

func (s *decodedStream) Read(buf []byte, offset int64) (n int, err error) {
	if offset < s.current {
		// TODO(jbd): Seek back if r.in is a ReadSeeker.
		return 0, errors.New("cannot seek back")
	}

	i := s.Info()
	length := i.Channels * i.BufferDepth * max

	// if buffer contains at least length number of frames, return them.
	if (offset-current)+length <= len(s.buf) {
		return s.buf[offset-current : (offset-current)+length]
	}

	// if buffer is not filled with enough data, read more from
	// s.in. If EOF from s.io, return EOF.
	panic("not yet implemented")
}

func (s *decodedStream) Info() (FrameInfo, error) {
	return FrameInfo{2, 16, 44000}, nil
}
