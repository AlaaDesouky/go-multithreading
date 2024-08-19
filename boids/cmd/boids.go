package boids

import (
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
	avgPosition, avgVelocity, separation := Vector2D{x: 0, y: 0}, Vector2D{x: 0, y: 0}, Vector2D{x: 0, y: 0}
	count := 0.0

	rwMu.RLock()
	for i := max(lower.x, 0); i <= min(upper.x, screenWidth); i++ {
		for j := max(lower.y, 0); j <= min(upper.y, screenHeight); j++ {
			if otherBoidID := boidMap[int(i)][int(j)]; otherBoidID != -1 && otherBoidID != b.id {
				if dist := boids[otherBoidID].position.Distance(b.position); dist < viewRadius {
					count++
					avgPosition = avgPosition.Add(boids[otherBoidID].position)
					avgVelocity = avgVelocity.Add(boids[otherBoidID].velocity)
					separation = separation.Add(b.position.Subtract(boids[otherBoidID].position).DivisionV(dist))
				}
			}

		}
	}
	rwMu.RUnlock()

	accel := Vector2D{x: b.borderBounce(b.position.x, screenWidth), y: b.borderBounce(b.position.y, screenHeight)}
	if count > 0 {
		avgPosition, avgVelocity = avgPosition.DivisionV(count), avgVelocity.DivisionV(count)

		accelCohesion := avgPosition.Subtract(b.position).MultiplyV(adjRate)
		accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyV(adjRate)
		accelSeparation := separation.MultiplyV(adjRate)


		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}

	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos - viewRadius {
		return 1 / (pos - maxBorderPos)
	}

	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()
	rwMu.Lock()
	b.velocity = b.velocity.Add(acceleration).Limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	rwMu.Unlock()
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		position: Vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		velocity: Vector2D{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
		id: bid,
	}

	boids[bid] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	go b.start()
}