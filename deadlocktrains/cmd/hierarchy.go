package deadlocktrains

import (
	"sort"
	"time"
)

func lockIntersectionInDistance(id, reserveStart, reserveEnd int, crossings []*Crossing) {
	var intersectionsToLock []*Intersection

	for _, crossing := range crossings {
		if reserveEnd >= crossing.Position && reserveStart <= crossing.Position && crossing.Intersection.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}

	sort.Slice(intersectionsToLock, func(i, j int) bool {
		return intersectionsToLock[i].ID < intersectionsToLock[j].ID
	})

	for _, intersection := range intersectionsToLock {
		intersection.Mutex.Lock()
		intersection.LockedBy = id
		time.Sleep(time.Millisecond * 10)
	}

}

func MoveTrainHierarchy(train *Train, distance int, crossings []*Crossing) {
	for train.Front < distance {
		train.Front += 1

		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				lockIntersectionInDistance(train.ID, crossing.Position, crossing.Position+train.TrainLength, crossings)
			}

			back := train.Front - train.TrainLength
			if back == crossing.Position {
				crossing.Intersection.LockedBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
		}

		time.Sleep(time.Millisecond * 30)
	}
}
