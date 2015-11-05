package audio_test

import "github.com/rakyll/experimental/audio"

func Example() {
	var s audio.Stream
	d := audio.NewMP3Decoder(s)
	d2 := audio.NewDecoder(d.Output())
	d3 := audio.NewDecoder(d2.Output())

	// decoding begins when d3.Ouput needs to be consumed.
	p, err := audio.NewPlayer(d3.Output())
	if err != nil {
		panic(err)
	}
	p.Play()
}
