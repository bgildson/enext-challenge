package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"

	"github.com/bgildson/enext-challenge/parser"
	"github.com/bgildson/enext-challenge/report"
)

func main() {
	gamesJSONPath := flag.String("games-json-path", "./games.json", "specify the json path of the parsed games.log")
	general := flag.Bool("general", true, "specify if the report should be general")
	flag.Parse()

	// try read source games
	data, err := ioutil.ReadFile(*gamesJSONPath)
	if err != nil {
		log.Fatalf("could not read source file: %v", err)
	}

	var gamesMap map[string]*parser.Game

	// deserialize games from source
	err = json.Unmarshal(data, &gamesMap)
	if err != nil {
		log.Fatalf("could not load games from source: %v", err)
	}

	// parse map for list
	var games []*parser.Game
	for _, g := range gamesMap {
		games = append(games, g)
	}

	// order ranking by id
	sort.Slice(games, func(i, j int) bool {
		x, _ := strconv.Atoi(games[i].ID)
		y, _ := strconv.Atoi(games[j].ID)
		return x < y
	})

	// print result, grouped or not based in general flag
	if *general {
		fmt.Println(report.ForGames(games))
	} else {
		for _, g := range games {
			fmt.Printf("%s\n\n", report.ForGame(g))
		}
	}
}
