# Multithreading In Go

Multi-threading examples, including a boids simulation in Go Lang

To run any of the code examples/scenarios, use the `-service` flag to specify the service type. The available service types are:

- `sync`
- `boids`
- `filesearch` `ARGS=ROOT,FILENAME`

### Run Using `go run` or `make`

Navigate to the project directory and use the following command to run a specific service:

```bash
cd ./go-multithreading
# using go run
go run main.go -service=<SERVICE_TYPE> -<args>=<ARGS>

# using make
make run-<SERVICE_TYPE> <ARGS>
```
