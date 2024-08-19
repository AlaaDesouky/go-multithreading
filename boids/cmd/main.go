package boids

import (
	"image/color"
	"log"
	"math/rand"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	title = "Boids in a box"

	screenWidth, screenHeight = 640, 460
	boidCount                 = 1000
	viewRadius = 13
	adjRate = 0.015
)

var (
	boids [boidCount]*Boid
	boidMap [screenWidth + 1][screenHeight + 1]int
	rwMu = sync.RWMutex{}
)

func getRandColorValue() uint8 {
	return uint8(rand.Intn(256))
}

func getColor() color.RGBA {
	return color.RGBA{getRandColorValue(), getRandColorValue(), getRandColorValue(), 255}
}

type Game struct{}

func (g Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		color := getColor()
		screen.Set(int(boid.position.x+1), int(boid.position.y), color)
		screen.Set(int(boid.position.x-1), int(boid.position.y), color)
		screen.Set(int(boid.position.x), int(boid.position.y-1), color)
		screen.Set(int(boid.position.x), int(boid.position.y+1), color)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func Run() {
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	for i := range boidCount {
		createBoid(i)
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle(title)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}