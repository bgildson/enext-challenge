package report

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/bgildson/enext-challenge/parser"
)

// Player represents a game player
type Player struct {
	Name   string
	Points int
}

// NewPlayer creates a new Player instance
func NewPlayer(name string, points int) *Player {
	return &Player{
		Name:   name,
		Points: points,
	}
}

// Ranking accumulates the player points
type Ranking struct {
	TotalKills int
	Players    map[string]*Player
}

// NewRanking creates a new Ranking instance
func NewRanking() *Ranking {
	return &Ranking{
		TotalKills: 0,
		Players:    map[string]*Player{},
	}
}

// AddGame integrate game points to the ranking
func (r *Ranking) AddGame(g *parser.Game) {
	r.TotalKills += g.TotalKills
	for p, k := range g.Kills {
		if _, ok := r.Players[p]; !ok {
			r.Players[p] = NewPlayer(p, 0)
		}
		r.Players[p].Points += k
	}
}

// Ordered returns the players ordered by points
func (r *Ranking) Ordered() []*Player {
	var p []*Player
	for _, v := range r.Players {
		p = append(p, v)
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].Points > p[j].Points || (p[i].Points == p[j].Points && p[i].Name < p[j].Name)
	})

	return p
}

// Report generates a text for the ranking
func (r *Ranking) Report() string {
	header := `Position | Player                         | Points`
	body := ""
	for i, p := range r.Ordered() {
		namePadLeft := strings.Repeat(" ", int(math.Max(0, float64(30-len(p.Name)))))
		body += fmt.Sprintf("%8d | %s%s | %d\n", i+1, p.Name, namePadLeft, p.Points)
	}
	body = strings.TrimRight(body, "\n")
	return fmt.Sprintf("%s\n%s", header, body)
}

// ForGame generates a ranking for one game
func ForGame(g *parser.Game) string {
	gameHeader := fmt.Sprintf("Game %s", g.ID)
	totalKillsHeader := fmt.Sprintf("Total Kills: %d", g.TotalKills)

	headerFormat := fmt.Sprintf("%%s%%%ds", 50-len(gameHeader))

	header := fmt.Sprintf(headerFormat, gameHeader, totalKillsHeader)

	r := NewRanking()
	r.AddGame(g)

	body := r.Report()

	return fmt.Sprintf("%s\n%s", header, body)
}

// ForGames generates a ranking for many games
func ForGames(gs []*parser.Game) string {
	r := NewRanking()
	for _, g := range gs {
		r.AddGame(g)
	}

	gameHeader := fmt.Sprintf("General Ranking")
	totalKillsHeader := fmt.Sprintf("Total Kills: %d", r.TotalKills)

	headerFormat := fmt.Sprintf("%%s%%%ds", 50-len(gameHeader))

	header := fmt.Sprintf(headerFormat, gameHeader, totalKillsHeader)

	body := r.Report()

	return fmt.Sprintf("%s\n%s", header, body)
}
