package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type SERVICE_TYPE string

var (
	SYNC SERVICE_TYPE = "sync"
)

func main() {
	services := []SERVICE_TYPE{SYNC}

	serviceType := flag.String("service", "", fmt.Sprintf("Specify the service to run: %v", services))

	flag.Parse()

	if *serviceType == "" {
		fmt.Println("Please specify a service to run using the -service flag.")
		os.Exit(1)
	}

	switch SERVICE_TYPE(*serviceType) {
	case SYNC:
		runService(SERVICE_TYPE(*serviceType))
	default:
		fmt.Printf("Unknown service type: %s\n", *serviceType)
		os.Exit(1)
	}
}

func runService(service SERVICE_TYPE) {
	fmt.Printf("Running service: %s\n", service)
	cmd := exec.Command("go", "run", fmt.Sprintf("./%s/main.go", service))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to run %s service: %v\n", service, err)
		os.Exit(1)
	}
}
