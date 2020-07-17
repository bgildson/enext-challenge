package repository

import (
	"reflect"
	"testing"

	"github.com/bgildson/enext-challenge/api/database"
	"github.com/bgildson/enext-challenge/parser"
)

func TestGamesRepository(t *testing.T) {
	games := []*parser.Game{
		{
			ID:         "1",
			TotalKills: 0,
			Players:    []string{},
			Kills:      map[string]int{},
		},
		{
			ID:         "2",
			TotalKills: 11,
			Players:    []string{"Isgalamido", "Mocinha"},
			Kills: map[string]int{
				"Isgalamido": -7,
				"Mocinha":    0,
			},
		},
	}
	game := parser.Game{
		ID:         "2",
		TotalKills: 11,
		Players:    []string{"Isgalamido", "Mocinha"},
		Kills: map[string]int{
			"Isgalamido": -7,
			"Mocinha":    0,
		},
	}
	dbSuccess := database.NewMockDatabase(
		func() ([]map[string]interface{}, error) {
			return []map[string]interface{}{
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
			}, nil
		},
		func(id string) (map[string]interface{}, error) {
			return map[string]interface{}{
				"id":          "2",
				"total_kills": 11,
				"players":     []string{"Isgalamido", "Mocinha"},
				"kills": map[string]int{
					"Isgalamido": -7,
					"Mocinha":    0,
				},
			}, nil
		},
	)
	dbFailure := database.NewMockDatabase(
		func() ([]map[string]interface{}, error) {
			return nil, database.ErrCouldNotDeserializeDatabaseContent
		},
		func(id string) (map[string]interface{}, error) {
			return nil, database.ErrGameNotFound
		},
	)
	repoSuccess := NewJSONGamesRepository(dbSuccess)
	repoFailure := NewJSONGamesRepository(dbFailure)

	t.Run("GetAll", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			r, err := repoSuccess.GetAll()
			if err != nil {
				t.Errorf("could not get games: %v", err)
			}

			if !reflect.DeepEqual(r, games) {
				t.Errorf("was expecting\n%v\nbut returns\n%v\n", games, r)
			}
		})

		t.Run("failure", func(t *testing.T) {
			r, err := repoFailure.GetAll()
			if err == nil {
				t.Errorf("was expecting a handled error, but was not catched")
			}

			if r != nil {
				t.Errorf("was expecting\n%v\nbut returns\n%v\n", nil, r)
			}
		})
	})

	t.Run("GetByID", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			r, err := repoSuccess.GetByID(game.ID)
			if err != nil {
				t.Errorf("could not get game %v: %v", game.ID, err)
			}

			if !reflect.DeepEqual(*r, game) {
				t.Errorf("was expecting\n%v\nbut returns\n%v\n", game, r)
			}
		})

		t.Run("failure", func(t *testing.T) {
			gameID := "-1"
			r, err := repoFailure.GetByID(gameID)
			if err == nil {
				t.Errorf("was expecting a handled error, but was not catched")
			}

			if r != nil {
				t.Errorf("was expecting\n%v\nbut returns\n%v\n", nil, r)
			}
		})
	})
}
