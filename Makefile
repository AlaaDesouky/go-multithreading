# Define variables
BINARY_NAME=bin
SERVICE_TYPE=
ROOT=
FILENAME=

# Default target
all: build

# Build the application
build:
	go build -o $(BINARY_NAME) main.go

# Run the application with a specific service type
run: build
	@echo "Running $(BINARY_NAME) with service type $(SERVICE_TYPE)"
	./$(BINARY_NAME) -service=$(SERVICE_TYPE) $(if $(ROOT),-root=$(ROOT)) $(if $(FILENAME),-filename=$(FILENAME))

# Clean up generated files
clean:
	@echo "Cleaning up..."
	@if exist $(BINARY_NAME) (del /q $(BINARY_NAME)) else (rm -f $(BINARY_NAME))

run-sync: SERVICE_TYPE=sync
run-sync: run

run-boids: SERVICE_TYPE=boids
run-boids: run

run-filesearch: SERVICE_TYPE=filesearch
run-filesearch: run

run-winddirection: SERVICE_TYPE=winddirection
run-winddirection: run

run-threadpool: SERVICE_TYPE=threadpool
run-threadpool: run

run-matrixmultiplication: SERVICE_TYPE=matrixmultiplication run
run-matrixmultiplication: run

.PHONY: all build run clean run-sync run-boids run-filesearch run-winddirection run-threadpool run-matrixmultiplication