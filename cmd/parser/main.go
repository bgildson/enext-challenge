package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/bgildson/enext-challenge/parser"
)

func main() {
	logPath := flag.String("log", "./games.log", "path to the log file")
	outPath := flag.String("out", "./games.json", "path to save processed log")
	flag.Parse()

	// open log file
	f, err := os.Open(*logPath)
	if err != nil {
		log.Fatalf("could not read the log file: %v", err)
	}

	// capture log lines
	lines := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	// process log and generate games for every log
	games := parser.ProcessLines(lines)

	// generate a map to add "ids" for every game
	gamesMap := map[string]*parser.Game{}
	for i, game := range games {
		gameKey := strconv.Itoa(i + 1)
		gamesMap[gameKey] = game
	}

	// serialize games map
	gamesJSON, err := json.MarshalIndent(gamesMap, "", "  ")
	if err != nil {
		log.Fatalf("could not serialize games: %v", err)
	}

	// write serialized data
	if err := ioutil.WriteFile(*outPath, gamesJSON, 0644); err != nil {
		log.Fatalf("could not write serialized games: %v", err)
	}
}
