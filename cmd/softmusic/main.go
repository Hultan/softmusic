package main

import (
	mp3 "github.com/hultan/softmusic/internal/mp3"
	"log"
)

func runMP3() error {
	player := mp3.NewSoftMusicMP3()
	player.PlayMP3()

	return nil
}

func main() {
	if err := runMP3(); err != nil {
		log.Fatal(err)
	}
}