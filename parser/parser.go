package parser

import (
	"regexp"
	"strconv"
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

// NewLine creates a new Line instance
func NewLine(text string) *Line {
	return &Line{
		line: text,
	}
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
	ID         string         `json:"id"`
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

// NewGameEmpty creates a new Game instance
func NewGameEmpty() *Game {
	return &Game{
		ID:         "",
		TotalKills: 0,
		Players:    []string{},
		Kills:      map[string]int{},
	}
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

	// when the killer and the dead was the same player, does nothing
	if k.Killer == k.Dead {
		return
	}

	// contabilize new kill to player
	if k.Killer == "<world>" {
		g.Kills[k.Dead]--
	} else {
		g.Kills[k.Killer]++
	}
}

// ProcessLines takes as input the log lines, process it and return stat games
func ProcessLines(lines []string) []*Game {
	var gs []*Game

	var g *Game
	currentID := 1
	for _, text := range lines {
		l := NewLine(text)

		if l.IsStartGame() {
			if g != nil {
				gs = append(gs, g)
			}

			g = NewGameEmpty()
			g.ID = strconv.Itoa(currentID)
			currentID++
		}

		k := l.AsKill()
		if k != nil {
			g.AddKill(k)
		}
	}
	if g != nil {
		gs = append(gs, g)
	}

	return gs
}
