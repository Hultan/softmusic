package softmusic

import "github.com/hultan/softmusic/internal/softmusic/internal/mp3"

type SoftMusic struct {

}

func NewSoftMusic() *SoftMusic {
	return new(SoftMusic)
}

func (s *SoftMusic) Play() {
	player := mp3.NewSoftMusicMP3()
	player.PlayMP3()
}