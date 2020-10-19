package softmusic

import "github.com/hultan/softmusic/pkg/mp3"

type SoftMusic struct {

}

func NewSoftMusic() *SoftMusic {
	return new(SoftMusic)
}

func (s *SoftMusic) Play() {
	player := mp3.NewMP3()
	player.PlayMP3()
}