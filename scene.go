package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	count         int
	entityManager EntityManager
	camera        Camera
}

func InitGameScene(entitySettings EntitySettings) *Game {
	return &Game{0, InitEntityManager(entitySettings), InitCamera(0, 0)}
}

func (g *Game) Update() error {
	g.count++
	g.camera.Update()

	// pass entity manager through system to update
	PeopleMovement(&g.entityManager)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.camera.Draw(screen, &g.entityManager, g.count)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\n", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
