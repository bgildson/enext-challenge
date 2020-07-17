package report

import (
	"reflect"
	"testing"

	"github.com/bgildson/enext-challenge/parser"
)

func TestRankingAddGame(t *testing.T) {
	tt := []struct {
		description string
		in          []*parser.Game
		out         *Ranking
	}{
		{
			description: "only for one game",
			in: []*parser.Game{
				{
					ID:         "1",
					TotalKills: 5,
					Players:    []string{"player one", "player two"},
					Kills: map[string]int{
						"player one": 2,
						"player two": 3,
					},
				},
			},
			out: &Ranking{
				TotalKills: 5,
				Players: map[string]*Player{
					"player one": NewPlayer("player one", 2),
					"player two": NewPlayer("player two", 3),
				},
			},
		},
		{
			description: "for many games",
			in: []*parser.Game{
				{
					ID:         "1",
					TotalKills: 5,
					Players:    []string{"player one", "player two"},
					Kills: map[string]int{
						"player one": 2,
						"player two": 3,
					},
				},
				{
					ID:         "2",
					TotalKills: 4,
					Players:    []string{"player one", "player three"},
					Kills: map[string]int{
						"player one":   3,
						"player three": 1,
					},
				},
			},
			out: &Ranking{
				TotalKills: 9,
				Players: map[string]*Player{
					"player one":   NewPlayer("player one", 5),
					"player two":   NewPlayer("player two", 3),
					"player three": NewPlayer("player three", 1),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			r := NewRanking()
			for _, g := range tc.in {
				r.AddGame(g)
			}
			if !reflect.DeepEqual(r, tc.out) {
				t.Errorf("was expecting %#v, but returns %#v", tc.out, r)
			}
		})
	}
}

func TestRankingOrdered(t *testing.T) {
	tt := []struct {
		description string
		in          *Ranking
		out         []*Player
	}{
		{
			description: "all player with different points",
			in: &Ranking{
				Players: map[string]*Player{
					"player one": {
						Name:   "player one",
						Points: 2,
					},
					"player two": {
						Name:   "player two",
						Points: 1,
					},
					"player three": {
						Name:   "player three",
						Points: 5,
					},
				},
			},
			out: []*Player{
				NewPlayer("player three", 5),
				NewPlayer("player one", 2),
				NewPlayer("player two", 1),
			},
		},
		{
			description: "two players with the same points",
			in: &Ranking{
				Players: map[string]*Player{
					"player one": {
						Name:   "player one",
						Points: 2,
					},
					"player two": {
						Name:   "player two",
						Points: 1,
					},
					"player three": {
						Name:   "player three",
						Points: 1,
					},
				},
			},
			out: []*Player{
				NewPlayer("player one", 2),
				NewPlayer("player three", 1),
				NewPlayer("player two", 1),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if o := tc.in.Ordered(); !reflect.DeepEqual(o, tc.out) {
				t.Errorf("was expecting %#v, but returns %#v", tc.out, o)
			}
		})
	}
}

func TestRankingReport(t *testing.T) {
	tt := []struct {
		description string
		in          *Ranking
		out         string
	}{
		{
			description: "all player with different points",
			in: &Ranking{
				Players: map[string]*Player{
					"player one":   NewPlayer("player one", 1),
					"player two":   NewPlayer("player two", 5),
					"player three": NewPlayer("player three", 3),
				},
			},
			out: `Position | Player                         | Points
       1 | player two                     | 5
       2 | player three                   | 3
       3 | player one                     | 1`,
		},
		{
			description: "two players with the same points",
			in: &Ranking{
				Players: map[string]*Player{
					"player one":   NewPlayer("player one", 1),
					"player two":   NewPlayer("player two", 3),
					"player three": NewPlayer("player three", 3),
				},
			},
			out: `Position | Player                         | Points
       1 | player three                   | 3
       2 | player two                     | 3
       3 | player one                     | 1`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if r := tc.in.Report(); r != tc.out {
				t.Errorf("\nwas expecting\n%v\nbut receives\n%v", tc.out, r)
			}
		})
	}
}

func TestForGame(t *testing.T) {
	tt := []struct {
		description string
		in          *parser.Game
		out         string
	}{
		{
			description: "a normal game",
			in: &parser.Game{
				ID:         "1",
				TotalKills: 5,
				Players:    []string{"player one", "player two"},
				Kills: map[string]int{
					"player one": 2,
					"player two": 3,
				},
			},
			out: `Game 1                              Total Kills: 5
Position | Player                         | Points
       1 | player two                     | 3
       2 | player one                     | 2`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if r := ForGame(tc.in); r != tc.out {
				t.Errorf("was expecting\n%v\nbut returns\n%v", tc.out, r)
			}
		})
	}
}

func TestForGames(t *testing.T) {
	tt := []struct {
		description string
		in          []*parser.Game
		out         string
	}{
		{
			description: "a normal game",
			in: []*parser.Game{
				{
					ID:         "1",
					TotalKills: 5,
					Players:    []string{"player one", "player two"},
					Kills: map[string]int{
						"player one": 2,
						"player two": 3,
					},
				},
				{
					ID:         "2",
					TotalKills: 7,
					Players:    []string{"player one", "player two"},
					Kills: map[string]int{
						"player one":   2,
						"player three": 2,
					},
				},
			},
			out: `General Ranking                    Total Kills: 12
Position | Player                         | Points
       1 | player one                     | 4
       2 | player two                     | 3
       3 | player three                   | 2`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if r := ForGames(tc.in); r != tc.out {
				t.Errorf("was expecting\n%v\nbut returns\n%v", tc.out, r)
			}
		})
	}
}
