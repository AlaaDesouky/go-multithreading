package deadlocktrains

import "time"

func MoveTrainDeadlock(train *Train, distance int, crossings []*Crossing) {
	for train.Front < distance {
		train.Front += 1

		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				crossing.Intersection.Mutex.Lock()
				crossing.Intersection.LockedBy = train.ID
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
