package parser

import (
	"reflect"
	"testing"
)

func TestLineAsKill(t *testing.T) {
	tt := []struct {
		in  string
		out *Kill
	}{
		{
			in: "21:42 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
			out: &Kill{
				Killer: "<world>",
				Dead:   "Isgalamido",
			},
		},
		{
			in: "22:06 Kill: 2 3 7: Isgalamido killed Mocinha by MOD_ROCKET_SPLASH",
			out: &Kill{
				Killer: "Isgalamido",
				Dead:   "Mocinha",
			},
		},
		{
			in:  "22:04 Item: 2 ammo_rockets",
			out: nil,
		},
		{
			in:  "22:11 ClientDisconnect: 3",
			out: nil,
		},
		{
			in:  "22:26 Item: 2 weapon_rocketlauncher",
			out: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			l := Line{tc.in}
			if r := l.AsKill(); !reflect.DeepEqual(r, tc.out) {
				t.Errorf("was expecting %v, but returns %v", tc.out, r)
			}
		})
	}
}

func TestLineIsStartGame(t *testing.T) {
	tt := []struct {
		in  string
		out bool
	}{
		{
			in:  `  0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`,
			out: true,
		},
		{
			in:  `  1:47 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\bot_minplayers\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`,
			out: true,
		},
		{
			in:  `  1:47 ClientConnect: 2`,
			out: false,
		},
		{
			in:  `  1:47 ClientUserinfoChanged: 2 n\Dono da Bola\t\0\model\sarge\hmodel\sarge\g_redteam\\g_blueteam\\c1\4\c2\5\hc\95\w\0\l\0\tt\0\tl\0`,
			out: false,
		},
		{
			in:  `  2:00 Kill: 1022 3 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`,
			out: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			l := Line{tc.in}
			if r := l.IsStartGame(); r != tc.out {
				t.Errorf("was expecting %v, but returns %v", tc.out, r)
			}
		})
	}
}

func TestGamePlayerExists(t *testing.T) {
	type In struct {
		game   Game
		player string
	}
	tt := []struct {
		description string
		in          In
		out         bool
	}{
		{
			description: "verify if exists a existing player",
			in: In{
				game: Game{
					TotalKills: 0,
					Players:    []string{"player one"},
					Kills: map[string]int{
						"player one": 0,
					},
				},
				player: "player one",
			},
			out: true,
		},
		{
			description: "verify if exists a non-existing player",
			in: In{
				game: Game{
					TotalKills: 0,
					Players:    []string{"player one"},
					Kills: map[string]int{
						"player one": 0,
					},
				},
				player: "player two",
			},
			out: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if r := tc.in.game.PlayerExists(tc.in.player); r != tc.out {
				t.Errorf("was expecting %v, but returns %v", tc.out, r)
			}
		})
	}
}

func TestGameAddPlayer(t *testing.T) {
	type Entry struct {
		game   Game
		player string
	}
	tt := []struct {
		description string
		in          Entry
		out         Game
	}{
		{
			description: "try add a non-existing player",
			in: Entry{
				game: Game{
					TotalKills: 0,
					Players:    []string{"player one"},
					Kills: map[string]int{
						"player one": 0,
					},
				},
				player: "player two",
			},
			out: Game{
				TotalKills: 0,
				Players:    []string{"player one", "player two"},
				Kills: map[string]int{
					"player one": 0,
					"player two": 0,
				},
			},
		},
		{
			description: "try add a existing player",
			in: Entry{
				game: Game{
					TotalKills: 0,
					Players:    []string{"player one"},
					Kills: map[string]int{
						"player one": 0,
					},
				},
				player: "player one",
			},
			out: Game{
				TotalKills: 0,
				Players:    []string{"player one"},
				Kills: map[string]int{
					"player one": 0,
				},
			},
		},
		{
			description: "try add <world> as player",
			in: Entry{
				game: Game{
					TotalKills: 0,
					Players:    []string{"player one"},
					Kills: map[string]int{
						"player one": 0,
					},
				},
				player: "<world>",
			},
			out: Game{
				TotalKills: 0,
				Players:    []string{"player one"},
				Kills: map[string]int{
					"player one": 0,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			g := tc.in.game
			g.AddPlayer(tc.in.player)
			if !reflect.DeepEqual(g, tc.out) {
				t.Errorf("was expecting %v, but returns %v", tc.out, g)
			}
		})
	}
}

func TestGameAddKill(t *testing.T) {
	type Entry struct {
		game Game
		kill Kill
	}
	tt := []struct {
		description string
		in          Entry
		out         Game
	}{
		{
			description: "add a kill from player to player",
			in: Entry{
				game: Game{
					TotalKills: 4,
					Players:    []string{"player one", "player two"},
					Kills: map[string]int{
						"player one": 2,
						"player two": 2,
					},
				},
				kill: Kill{
					Killer: "player one",
					Dead:   "player two",
				},
			},
			out: Game{
				TotalKills: 5,
				Players:    []string{"player one", "player two"},
				Kills: map[string]int{
					"player one": 3,
					"player two": 2,
				},
			},
		},
		{
			description: "add a kill from world to player",
			in: Entry{
				game: Game{
					TotalKills: 4,
					Players:    []string{"player one", "player two"},
					Kills: map[string]int{
						"player one": 2,
						"player two": 2,
					},
				},
				kill: Kill{
					Killer: "<world>",
					Dead:   "player two",
				},
			},
			out: Game{
				TotalKills: 5,
				Players:    []string{"player one", "player two"},
				Kills: map[string]int{
					"player one": 2,
					"player two": 1,
				},
			},
		},
		{
			description: "add a kill with the killer been a new player",
			in: Entry{
				game: Game{
					TotalKills: 4,
					Players:    []string{"player one", "player two"},
					Kills: map[string]int{
						"player one": 2,
						"player two": 2,
					},
				},
				kill: Kill{
					Killer: "player three",
					Dead:   "player two",
				},
			},
			out: Game{
				TotalKills: 5,
				Players:    []string{"player one", "player two", "player three"},
				Kills: map[string]int{
					"player one":   2,
					"player two":   2,
					"player three": 1,
				},
			},
		},
		{
			description: "add a kill with the dead been a new player",
			in: Entry{
				game: Game{
					TotalKills: 4,
					Players:    []string{"player one", "player two"},
					Kills: map[string]int{
						"player one": 2,
						"player two": 2,
					},
				},
				kill: Kill{
					Killer: "player two",
					Dead:   "player three",
				},
			},
			out: Game{
				TotalKills: 5,
				Players:    []string{"player one", "player two", "player three"},
				Kills: map[string]int{
					"player one":   2,
					"player two":   3,
					"player three": 0,
				},
			},
		},
		{
			description: "add a kill with killer and dead been a new player",
			in: Entry{
				game: Game{
					TotalKills: 4,
					Players:    []string{"player one", "player two"},
					Kills: map[string]int{
						"player one": 2,
						"player two": 2,
					},
				},
				kill: Kill{
					Killer: "player three",
					Dead:   "player four",
				},
			},
			out: Game{
				TotalKills: 5,
				Players:    []string{"player one", "player two", "player three", "player four"},
				Kills: map[string]int{
					"player one":   2,
					"player two":   2,
					"player three": 1,
					"player four":  0,
				},
			},
		},
		{
			description: "add a kill with killer and dead been the same player",
			in: Entry{
				game: Game{
					TotalKills: 4,
					Players:    []string{"player one", "player two"},
					Kills: map[string]int{
						"player one": 2,
						"player two": 2,
					},
				},
				kill: Kill{
					Killer: "player three",
					Dead:   "player three",
				},
			},
			out: Game{
				TotalKills: 5,
				Players:    []string{"player one", "player two", "player three"},
				Kills: map[string]int{
					"player one":   2,
					"player two":   2,
					"player three": 0,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			g := tc.in.game
			g.AddKill(&tc.in.kill)
			if !reflect.DeepEqual(g, tc.out) {
				t.Errorf("was expecting %v, but returns %v", tc.out, g)
			}
		})
	}
}
