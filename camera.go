package main

import "github.com/hajimehoshi/ebiten/v2"

type Vec2D struct {
	X int
	Y int
}

type Camera struct {
	Position Vec2D
}

func InitCamera(x int, y int) Camera {
	return Camera{Vec2D{x, y}}
}

func (c *Camera) Draw(screen *ebiten.Image, manager *EntityManager, count int) {
	// people
	averageGroupWealth := manager.World.WorldWealth / float64(len(manager.GroupPeople))
	for _, population := range manager.GroupPeople {
		imgResizeFactor := population.GroupWealth / averageGroupWealth
		img, drawOp := population.DrawParameter(screen, count)
		drawOp.GeoM.Translate(-float64(c.Position.X), -float64(c.Position.Y))
		drawOp.GeoM.Scale(imgResizeFactor, imgResizeFactor)
		screen.DrawImage(img, drawOp)
	}

	// market
	img, drawOp := manager.Market.DrawParameter(screen)
	drawOp.GeoM.Translate(-float64(c.Position.X), -float64(c.Position.Y))
	screen.DrawImage(img, drawOp)

	// .....

	// .....

}

func (c *Camera) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		c.Position.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		c.Position.X += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		c.Position.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		c.Position.Y += 1
	}
	return nil
}
