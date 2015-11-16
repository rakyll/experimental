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

func (s *decodedStream) Read(offset int64, max int64) (buf []byte, n int, err error) {
	panic("not yet implemented")
}

func (s *decodedStream) Info() (FrameInfo, error) {
	return FrameInfo{2, 16, 44000}, nil
}
