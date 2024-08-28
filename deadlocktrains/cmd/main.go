package deadlocktrains

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	noOfCrossings = 4
	trainLength   = 70
)

var (
	trains        [noOfCrossings]*Train
	intersections [noOfCrossings]*Intersection

	windowSize  = 320
	windowTitle = "Train in a box"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawTracks(screen)
	DrawIntersections(screen)
	DrawTrains(screen)
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return windowSize, windowSize
}

func RunDeadlock() {
	for i := 0; i < noOfCrossings; i++ {
		trains[i] = &Train{ID: i, TrainLength: trainLength, Front: 0}
		intersections[i] = &Intersection{ID: i, Mutex: sync.Mutex{}, LockedBy: -1}
	}

	// Deadlock
	go MoveTrainDeadlock(trains[0], 300, []*Crossing{{Position: 125, Intersection: intersections[0]},
		{Position: 175, Intersection: intersections[1]}})

	go MoveTrainDeadlock(trains[1], 300, []*Crossing{{Position: 125, Intersection: intersections[1]},
		{Position: 175, Intersection: intersections[2]}})

	go MoveTrainDeadlock(trains[2], 300, []*Crossing{{Position: 125, Intersection: intersections[2]},
		{Position: 175, Intersection: intersections[3]}})

	go MoveTrainDeadlock(trains[3], 300, []*Crossing{{Position: 125, Intersection: intersections[3]},
		{Position: 175, Intersection: intersections[0]}})

	ebiten.SetWindowSize(windowSize*3, windowSize*3)
	ebiten.SetWindowTitle(windowTitle + " - Deadlock")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
