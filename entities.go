package main

import (
	"bytes"
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

type EntitySettings struct {
	Population int
}

type EntityManager struct {
	Population  []People
	GroupPeople []GroupPeople
	Market      Market
	World       World
}

func InitEntityManager(entitySettings EntitySettings) EntityManager {

	populationCount := entitySettings.Population
	peopleList := make([]People, populationCount)
	grounpPeople := []GroupPeople{}
	var currentGrounpId string
	count := 0
	for i := 0; i < populationCount; i++ {
		if i%10000 == 0 {
			log.Println(count)
			obj := InitGroupPeople()
			grounpPeople = append(grounpPeople, obj)
			currentGrounpId = obj.Id
			count += 1
		}
		peopleList[i] = InitPerson(currentGrounpId)
	}

	return EntityManager{
		Population:  peopleList,
		Market:      InitMarket(),
		World:       InitWorld(),
		GroupPeople: grounpPeople,
	}
}

const (
	// describe the sub image size
	personFrameWidth  = 32
	personFrameHeight = 32
	peopleFrameCount  = 8

	// describes the sub image position
	personImgFrameX = 0
	personImgFrameY = 32

	// rescale image size
	personImgScaleFactor = 0.8
	// these two descibes the center position of the bbox drawing on the scene
	rawDrawCenterX = -float64(personFrameWidth) / 2 * personImgScaleFactor
	rawDrawCenterY = -float64(personFrameHeight) / 2 * personImgScaleFactor
)

var PeopleImage *ebiten.Image

type GroupPeople struct {
	Id           string
	CenterX      float64
	CenterY      float64
	VelocityX    float64
	VelocityY    float64
	Acceleration float64
	ImgNo        int
	ImgDirection int
	GroupWealth  float64
}

func InitGroupPeople() GroupPeople {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if PeopleImage == nil {
		PeopleImage = ebiten.NewImageFromImage(img)
	}
	x, y := float64(r.Intn(ScreenWidth*2)), float64(r.Intn(ScreenHeight*2))
	if err != nil {
		log.Fatal(err)
	}
	return GroupPeople{
		Id:           uuid.New().String(),
		CenterX:      x,
		CenterY:      y,
		VelocityX:    r.Float64(),
		VelocityY:    r.Float64(),
		Acceleration: r.Float64(),
		ImgNo:        r.Intn(5),
		ImgDirection: r.Intn(2),
		GroupWealth:  0,
	}
}

func (d *GroupPeople) DrawParameter(screen *ebiten.Image, count int) (*ebiten.Image, *ebiten.DrawImageOptions) {
	// set image position
	imgX := rawDrawCenterX + d.CenterX
	imgY := rawDrawCenterY + d.CenterY

	op := &ebiten.DrawImageOptions{}
	if d.ImgDirection == 1 {
		op.GeoM.Scale(-1, 1)
	}
	op.GeoM.Scale(personImgScaleFactor, personImgScaleFactor)
	op.GeoM.Translate(imgX, imgY)
	// set which sub image
	i := (count + d.ImgNo/5) % peopleFrameCount
	sx, sy := personImgFrameX+i*personFrameWidth, personImgFrameY

	return PeopleImage.SubImage(image.Rect(sx, sy, sx+personFrameWidth, sy+personFrameHeight)).(*ebiten.Image), op
}

type People struct {
	GroupId      string
	Id           string
	Productivity float64
	Savings      float64
	Operations   []Operations
}

func InitPerson(grounpId string) People {
	return People{
		Id:           uuid.New().String(),
		Productivity: rand.Float64() * 500,
		Savings:      0,
		Operations:   []Operations{},
		GroupId:      grounpId,
	}
}

const (
	WorldWidth  = ScreenWidth
	WorldHeight = ScreenHeight
)

type World struct {
	Width       float64
	Height      float64
	WorldWealth float64
}

func InitWorld() World {
	return World{
		Width:       WorldWidth,
		Height:      WorldHeight,
		WorldWealth: 0,
	}
}

var MarketImage *ebiten.Image

type Market struct {
	CenterX float64
	CenterY float64
}

func InitMarket() Market {
	img, _, _ := ebitenutil.NewImageFromFile("./content/shop.png")
	if MarketImage == nil {
		MarketImage = img
	}
	return Market{
		CenterX: ScreenWidth / 2,
		CenterY: ScreenHeight / 2,
	}
}

func (m *Market) DrawParameter(screen *ebiten.Image) (*ebiten.Image, *ebiten.DrawImageOptions) {
	// set image position
	imgX := m.CenterX
	imgY := m.CenterY

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(imgX, imgY)
	return MarketImage, op
}
