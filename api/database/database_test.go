package database

import (
	"encoding/json"
	"os"
	"path"
	"reflect"
	"sort"
	"testing"
)

func TestNewJSONDatabase(t *testing.T) {
	basePath, err := os.Getwd()
	fixturesPath := path.Join(basePath, "..", "..", "fixtures")
	if err != nil {
		t.Errorf("could not determine where the app is running: %v", err)
	}
	tt := []struct {
		description string
		in          string
		out         error
	}{
		{
			description: "success create and load database data",
			in:          path.Join(fixturesPath, "games.json"),
			out:         nil,
		},
		{
			description: "problems when looking for the database file",
			in:          path.Join(fixturesPath, "games_nonexistent.json"),
			out:         ErrDatabaseFileNotFound,
		},
		{
			description: "problems ummarshal file content",
			in:          path.Join(fixturesPath, "games_malformed.json"),
			out:         ErrCouldNotDeserializeDatabaseContent,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			if _, err := NewJSONDatabase(tc.in); err != tc.out {
				t.Errorf(`was expecting "%v" error, but returns "%v" error`, tc.out, err)
			}
		})
	}
}

func TestJSONDatabaseGet(t *testing.T) {
	basePath, err := os.Getwd()
	if err != nil {
		t.Errorf("could not get database base url")
	}
	gamesPath := path.Join(basePath, "..", "..", "fixtures", "games.json")
	d, err := NewJSONDatabase(gamesPath)
	if err != nil {
		t.Errorf("could load fixture to database: %v", err)
	}
	tt := []struct {
		description string
		in          Database
		out         []map[string]interface{}
	}{
		{
			description: "reply correctly",
			in:          d,
			out: []map[string]interface{}{
				{
					"id":          "1",
					"total_kills": 0,
					"players":     []string{},
					"kills":       map[string]int{},
				},
				{
					"id":          "2",
					"total_kills": 11,
					"players":     []string{"Isgalamido", "Mocinha"},
					"kills": map[string]int{
						"Isgalamido": -7,
						"Mocinha":    0,
					},
				},
				{
					"id":          "3",
					"total_kills": 4,
					"players":     []string{"Isgalamido", "Mocinha", "Zeh", "Dono da Bola"},
					"kills": map[string]int{
						"Dono da Bola": -1,
						"Isgalamido":   1,
						"Mocinha":      0,
						"Zeh":          -2,
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			r, _ := tc.in.Get()
			// make ordenation and comparation with bytes to correctly compare
			sort.Slice(r, func(i, j int) bool {
				a, _ := r[i]["id"].(string)
				b, _ := r[j]["id"].(string)
				return a < b
			})
			a, _ := json.Marshal(r)
			b, _ := json.Marshal(tc.out)
			if !reflect.DeepEqual(a, b) {
				t.Errorf("was expecting\n%v\nbut returns\n%v\n", tc.out, r)
			}
		})
	}
}

func TestJSONDatabaseGetByID(t *testing.T) {
	basePath, err := os.Getwd()
	if err != nil {
		t.Errorf("could not get database base url")
	}
	gamesPath := path.Join(basePath, "..", "..", "fixtures", "games.json")
	d, err := NewJSONDatabase(gamesPath)
	if err != nil {
		t.Errorf("could load fixture to database: %v", err)
	}
	tt := []struct {
		description string
		in          string
		game        map[string]interface{}
		err         error
	}{
		{
			description: "get one registry correctly by id",
			in:          "2",
			game: map[string]interface{}{
				"id":          "2",
				"total_kills": 11,
				"players":     []string{"Isgalamido", "Mocinha"},
				"kills": map[string]int{
					"Isgalamido": -7,
					"Mocinha":    0,
				},
			},
			err: nil,
		},
		{
			description: "get nonexisting registry by id",
			in:          "-1",
			game:        nil,
			err:         ErrGameNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			g, err := d.GetByID(tc.in)
			a, _ := json.Marshal(g)
			b, _ := json.Marshal(tc.game)
			if !reflect.DeepEqual(a, b) || err != tc.err {
				t.Errorf("was expecting\n%v\n%v\nbut returns\n%v\n%v\n", tc.game, tc.err, g, err)
			}
		})
	}
}
