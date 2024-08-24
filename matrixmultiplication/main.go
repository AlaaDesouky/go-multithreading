package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	matrixSize   = 250
	noOfMatrices = 100
)

var (
	matrixA = [matrixSize][matrixSize]int{}
	matrixB = [matrixSize][matrixSize]int{}
	result  = [matrixSize][matrixSize]int{}

	rwMU = sync.RWMutex{}
	cond = sync.NewCond(rwMU.RLocker())
	wg   = sync.WaitGroup{}
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}
func workOutRow(row int) {
	rwMU.RLock()
	for {
		wg.Done()
		cond.Wait()

		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}

	}
}

func main() {
	fmt.Println("Working...")
	wg.Add(matrixSize)

	for row := 0; row < matrixSize; row++ {
		go workOutRow(row)
	}

	start := time.Now()
	for i := 0; i < noOfMatrices; i++ {
		wg.Wait()
		rwMU.Lock()

		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)

		wg.Add(matrixSize)
		rwMU.Unlock()
		cond.Broadcast()
	}

	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took %s\n", elapsed)
}
