package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches []string
	wg      = sync.WaitGroup{}
	mu = sync.Mutex{}
)

func fileSearch(root, filename string) {
	fmt.Println("Searching in: ", root)
	files, _ := os.ReadDir(root)

	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			mu.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			mu.Unlock()
		}

		if file.IsDir() {
			wg.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}

	wg.Done()
}

func main() {
	root := flag.String("root", "", "Root directory to search")
	filename := flag.String("filename", "", "Filename to search for")
	flag.Parse()

	if *root == "" || *filename == "" {
		fmt.Println("Please specify both root directory and filename using -root and -filename flags.")
		os.Exit(1)
	}

	wg.Add(1)
	go fileSearch(*root, *filename)
	wg.Wait()

	if len(matches) == 0 {
		fmt.Printf("No Match: %s is not found in directory: %s\n", *filename, *root)
		return
	}

	for _, file := range matches {
		fmt.Println("Matched: ", file)
	}
}