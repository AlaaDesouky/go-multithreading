package main

import (
	deadlocktrains "go-multithreading/deadlocktrains/cmd"
)

func main() {
	// deadlocktrains.RunDeadlock()
	// deadlocktrains.RunDeadlockWithHierarchy()
	deadlocktrains.RunDeadlockWithArbitrator()
}
