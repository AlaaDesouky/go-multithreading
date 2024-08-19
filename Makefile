# Define variables
BINARY_NAME=bin
SERVICE_TYPE=

# Default target
all: build

# Build the application
build:
	go build -o $(BINARY_NAME) main.go

# Run the application with a specific service type
run: build
	@echo "Running $(BINARY_NAME) with service type $(SERVICE_TYPE)"
	./$(BINARY_NAME) -service=$(SERVICE_TYPE)

# Clean up generated files
clean:
	@echo "Cleaning up..."
	@del /q $(BINARY_NAME) 2>nul || rm -f $(BINARY_NAME)

# Specific target to run the sync service
run-sync: SERVICE_TYPE=sync
run-sync: run

.PHONY: all build run clean check-build check-build-windows run-sync