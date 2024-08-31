package deadlocktrains

import (
	"sync"
	"time"
)

var (
	controller = sync.Mutex{}
	cond       = sync.NewCond(&controller)
)

func isAllFree(intersectionsToLock []*Intersection) bool {
	for _, intersection := range intersectionsToLock {
		if intersection.LockedBy >= 0 {
			return false
		}
	}

	return true
}

func lockIntersectionInDistanceArbitrator(id, reserveStart, reserveEnd int, crossings []*Crossing) {
	var intersectionsToLock []*Intersection
	for _, crossing := range crossings {
		if reserveEnd >= crossing.Position && reserveStart <= crossing.Position && crossing.Intersection.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}

	controller.Lock()
	if !isAllFree(intersectionsToLock) {
		cond.Wait()
	}

	for _, intersection := range intersectionsToLock {
		intersection.LockedBy = id
		time.Sleep(time.Millisecond * 10)
	}
	controller.Unlock()
}

func MoveTrainArbitrator(train *Train, distance int, crossings []*Crossing) {
	for train.Front < distance {
		train.Front += 1

		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				lockIntersectionInDistanceArbitrator(train.ID, crossing.Position, crossing.Position+train.TrainLength, crossings)
			}

			back := train.Front - train.TrainLength
			if back == crossing.Position {
				controller.Lock()
				crossing.Intersection.LockedBy = -1
				cond.Broadcast()
				controller.Unlock()
			}
		}

		time.Sleep(time.Millisecond * 30)
	}
}
