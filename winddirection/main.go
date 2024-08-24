package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	windRegex     = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
	tafValidation = regexp.MustCompile(`.*TAF.*`)
	comment       = regexp.MustCompile(`\w*#.*`)
	metarClose    = regexp.MustCompile(`.*=`)
	variableWind  = regexp.MustCompile(`.*VRB\d{2}KT`)
	validWind     = regexp.MustCompile(`\d{5}KT`)
	windDirOnly   = regexp.MustCompile(`(\d{3})\d{2}KT`)
	windDist      [8]int

	textCh  = make(chan string)
	metarCh = make(chan []string)
	windsCh = make(chan []string)
	distCh  = make(chan [8]int)
)

func parseToArray() {
	for text := range textCh {
		lines := strings.Split(text, "\n")
		metoarSlice := make([]string, 0, len(lines))
		metarStr := ""

		for _, line := range lines {
			if tafValidation.MatchString(line) {
				break
			}

			if !comment.MatchString(line) {
				metarStr += strings.Trim(line, " ")
			}

			if metarClose.MatchString(line) {
				metoarSlice = append(metoarSlice, metarStr)
				metarStr = ""
			}
		}

		metarCh <- metoarSlice
	}

	close(metarCh)
}

func extractWindDirection() {
	for metars := range metarCh {
		winds := make([]string, 0, len(metars))
		for _, metar := range metars {
			if windRegex.MatchString(metar) {
				winds = append(winds, windRegex.FindAllStringSubmatch(metar, -1)[0][1])
			}
		}

		windsCh <- winds
	}

	close(windsCh)
}

func minWindDistribution() {
	for winds := range windsCh {
		for _, wind := range winds {
			if variableWind.MatchString(wind) {
				for i := 0; i < 8; i++ {
					windDist[i]++
				}
			} else if validWind.MatchString(wind) {
				windStr := windDirOnly.FindAllStringSubmatch(wind, -1)[0][1]
				if d, err := strconv.ParseFloat(windStr, 64); err != nil {
					dirIdx := int(math.Round(d/45.0)) % 8
					windDist[dirIdx]++
				}
			}
		}
	}

	distCh <- windDist
	close(distCh)
}

func main() {
	//1. Change to array, each metar report is a separate item in the array
	go parseToArray()

	//2. Extract wind direction, EGLL 312350Z 07004KT CAVOK 12/09 Q1016 NOSIG= -> 070
	go extractWindDirection()

	//3. Assign to N, NE, E, SE, S, SW, W, NW, 070 -> E + 1
	go minWindDistribution()

	absPath, _ := filepath.Abs("./winddirection/metarfiles/")
	files, _ := os.ReadDir(absPath)

	start := time.Now()
	for _, file := range files {
		data, err := os.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			panic(err)
		}

		text := string(data)
		textCh <- text
	}
	close(textCh)

	results := <-distCh
	elapsed := time.Since(start)

	fmt.Printf("%v\n", results)
	fmt.Printf("Processing took %s\n", elapsed)
}
