package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type SERVICE_TYPE string

var (
	SYNC                  SERVICE_TYPE = "sync"
	BOIDS                 SERVICE_TYPE = "boids"
	FILE_SEARCH           SERVICE_TYPE = "filesearch"
	WIND_DIRECTION        SERVICE_TYPE = "winddirection"
	THREAD_POOL           SERVICE_TYPE = "threadpool"
	MATRIX_MULTIPLICATION SERVICE_TYPE = "matrixmultiplication"

	services = []SERVICE_TYPE{SYNC, BOIDS, FILE_SEARCH, WIND_DIRECTION, THREAD_POOL, MATRIX_MULTIPLICATION}
)

func main() {
	serviceType, otherFlags := parseFlags()

	if *serviceType == "" {
		fmt.Println("Please specify a service to run using the -service flag.")
		os.Exit(1)
	}

	switch SERVICE_TYPE(*serviceType) {
	case SYNC:
	case BOIDS:
	case FILE_SEARCH:
	case WIND_DIRECTION:
	case THREAD_POOL:
	case MATRIX_MULTIPLICATION:
		runService(SERVICE_TYPE(*serviceType), otherFlags)

	default:
		fmt.Printf("Unknown service type: %s\n", *serviceType)
		os.Exit(1)
	}
}

func parseFlags() (*string, []string) {
	serviceType := flag.String("service", "", fmt.Sprintf("Specify the service to run: %v", services))

	flag.String("root", "", "Specify the root directory for file search service")
	flag.String("filename", "", "Specify the filename to search for")

	flag.Parse()

	var otherFlags []string
	flag.VisitAll(func(f *flag.Flag) {
		if f.Name != "service" && f.Value.String() != "" {
			otherFlags = append(otherFlags, fmt.Sprintf("-%s=%s", f.Name, f.Value))
		}
	})

	return serviceType, otherFlags
}

func runService(service SERVICE_TYPE, otherFlags []string) {
	log := fmt.Sprintf("Running service: %s", service)

	if len(otherFlags) > 0 {
		log = fmt.Sprintf("%s with flags: %v", log, otherFlags)
	}

	fmt.Printf("%v\n", log)

	cmdArgs := append([]string{"run", fmt.Sprintf("./%s/main.go", service)}, otherFlags...)
	cmd := exec.Command("go", cmdArgs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to run %s service: %v\n", service, err)
		os.Exit(1)
	}
}
