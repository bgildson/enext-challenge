package main

import (
	"regexp"
)

// Kill represents a kill in game
type Kill struct {
	Killer string
	Dead   string
}

// Line works like a façade adapter improving cast for an line while abstract the implementation
type Line struct {
	line string
}

// AsKill try to handle the line as a kill
// when is not possible, returns nil
func (l *Line) AsKill() *Kill {
	re := regexp.MustCompile(`.*Kill:.*: (.*) killed (.*) by.*`)

	match := re.FindStringSubmatch(l.line)
	if match == nil {
		return nil
	}

	return &Kill{
		Killer: match[1],
		Dead:   match[2],
	}
}

// IsStartGame verify if the line indicates a starting new game
func (l *Line) IsStartGame() bool {
	m, _ := regexp.MatchString(`.* InitGame: .*`, l.line)
	return m
}

// Game represents the game stats
type Game struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

// PlayerExists verify if just exists a player
func (g *Game) PlayerExists(player string) bool {
	for _, p := range g.Players {
		if p == player {
			return true
		}
	}
	return false
}

// AddPlayer try add play, if just exists or is <world>, does nothing
func (g *Game) AddPlayer(player string) {
	if player == "<world>" || g.PlayerExists(player) {
		return
	}

	g.Players = append(g.Players, player)
	g.Kills[player] = 0
}

// AddKill handles a new kill insertion
func (g *Game) AddKill(k *Kill) {
	// increment new kill
	g.TotalKills++

	// try to add killer player
	g.AddPlayer(k.Killer)

	// try to add dead player
	g.AddPlayer(k.Dead)

	// contabilize new kill to player
	if k.Killer == "<world>" {
		g.Kills[k.Dead]--
	} else {
		g.Kills[k.Killer]++
	}
}
