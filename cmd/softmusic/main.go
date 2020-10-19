package main

import (
	"github.com/hultan/softmusic/internal/softmusic"
	"log"
)

func runSoftMusic() error {
	player := softmusic.NewSoftMusic()
	player.Play()

	return nil
}

func main() {
	if err := runSoftMusic(); err != nil {
		log.Fatal(err)
	}
}