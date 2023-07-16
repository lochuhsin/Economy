package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// macbook pro max size
	// generated using ebiten.ScreenSizeInFullscreen()
	// these values should be either placed in environment or ...etc
	ScreenWidth  = 1512
	ScreenHeight = 982
)

func main() {
	ScreenWidth, ScreenHeight := ebiten.ScreenSizeInFullscreen()
	log.Println(ScreenWidth, ScreenHeight)
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)

	// Initialize People with concurrency e.g 100 people per batch
	entitySettings := EntitySettings{
		Population: 1000000,
	}

	systemChannel := make(chan EntityManager, 100000)

	system := InitSystemManager(entitySettings)
	game := InitGameScene(systemChannel)

	go system.Run(systemChannel)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
