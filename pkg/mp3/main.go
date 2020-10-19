package mp3

import (
	"fmt"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"io"
	"os"
)

type SoftMusicMP3 struct {

}

func NewSoftMusicMP3() *SoftMusicMP3 {
	return new(SoftMusicMP3)
}

func (s *SoftMusicMP3) PlayMP3() error {
	f, err := os.Open("/home/per/code/softmusic/assets/2cellos.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	fmt.Printf("Length: %d[bytes]\n", d.Length())

	if _, err := io.Copy(p, d); err != nil {
		return err
	}
	return nil
}

