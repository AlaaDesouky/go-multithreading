package main

import (
	"flag"
	"fmt"
	deadlocktrains "go-multithreading/deadlocktrains/cmd"
)

type SIMULATION_TYPE string

var (
	DEADLOCK   SIMULATION_TYPE = "deadlock"
	HIERARCHY  SIMULATION_TYPE = "hierarchy"
	ARBITRATOR SIMULATION_TYPE = "arbitrator"
)

func main() {
	simulation := flag.String("simulation", "", "Simulation type to run")
	flag.Parse()

	fmt.Println(*simulation)

	switch SIMULATION_TYPE(*simulation) {
	case DEADLOCK:
		deadlocktrains.RunDeadlock()
	case HIERARCHY:
		deadlocktrains.RunDeadlockWithHierarchy()
	case ARBITRATOR:
		deadlocktrains.RunDeadlockWithArbitrator()
	default:
		deadlocktrains.RunDeadlock()
	}
}
