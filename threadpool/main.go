package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const numberOfThreads int = 8

var (
	polygonPointRegex = regexp.MustCompile(`\((\d*),(\d*)\)`)
	wg                = sync.WaitGroup{}
	inputCh           = make(chan string, 1000)
)

type Point2D struct {
	x int
	y int
}

func findArea() {
	for pointsStr := range inputCh {
		var points []Point2D

		for _, p := range polygonPointRegex.FindAllStringSubmatch(pointsStr, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])

			points = append(points, Point2D{x, y})
		}

		area := 0.0
		n := len(points)
		for i := 0; i < n; i++ {
			a, b := points[i], points[(i+1)%n]
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}

		fmt.Printf("Area: %v \n", math.Abs(area)/2.0)
	}

	wg.Done()
}

func main() {
	absPath, _ := filepath.Abs("./threadpool/")
	data, _ := os.ReadFile(filepath.Join(absPath, "polygons.txt"))
	text := string(data)

	for i := 0; i < numberOfThreads; i++ {
		go findArea()
	}

	wg.Add(numberOfThreads)
	start := time.Now()

	for _, line := range strings.Split(text, "\n") {
		inputCh <- line
	}

	close(inputCh)
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Processing took %s \n", elapsed)
}
