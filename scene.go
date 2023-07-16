package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	count            int
	camera           Camera
	systemChannel    chan EntityManager
	preEntityManager *EntityManager // save previous frame information incase it's blocked
}

func InitGameScene(systemChannel chan EntityManager) *Game {
	return &Game{0, InitCamera(0, 0), systemChannel, nil}
}

func (g *Game) Update() error {
	g.count++
	g.camera.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	select {
	case manager := <-g.systemChannel:
		g.preEntityManager = &manager
		g.camera.Draw(screen, &manager, g.count)
	default:
		if g.preEntityManager != nil {
			g.camera.Draw(screen, g.preEntityManager, g.count)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\n", ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
