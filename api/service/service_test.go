package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bgildson/enext-challenge/api/repository"
	"github.com/bgildson/enext-challenge/parser"
)

func TestGamesService(t *testing.T) {
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
	game := &parser.Game{
		ID:         "2",
		TotalKills: 11,
		Players:    []string{"Isgalamido", "Mocinha"},
		Kills: map[string]int{
			"Isgalamido": -7,
			"Mocinha":    0,
		},
	}
	repositorySuccess := repository.NewMockGamesRepository(
		func() ([]*parser.Game, error) {
			return games, nil
		},
		func(id string) (*parser.Game, error) {
			return game, nil
		},
	)
	repositoryFailure := repository.NewMockGamesRepository(
		func() ([]*parser.Game, error) {
			return nil, fmt.Errorf("occur an error")
		},
		func(id string) (*parser.Game, error) {
			return nil, fmt.Errorf("occur an error")
		},
	)
	serviceSuccess := NewGamesService(repositorySuccess)
	serviceFailure := NewGamesService(repositoryFailure)

	t.Run("List", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			r, err := serviceSuccess.List()
			if err != nil {
				t.Errorf("an unexpected error occurred: %v", err)
			}

			if !reflect.DeepEqual(r, games) {
				t.Errorf("was expecting\n%#v\nbut returns\n%#v\n", games, r)
			}
		})

		t.Run("failure", func(t *testing.T) {
			r, err := serviceFailure.List()
			if err == nil {
				t.Errorf(`was expecting an error, but returns: "%v"`, err)
			}

			if r != nil {
				t.Errorf("a result value in an error case should be nil, but returns: %v", r)
			}
		})
	})

	t.Run("Find", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			r, err := serviceSuccess.Find(game.ID)
			if err != nil {
				t.Errorf("an unexpected error occurred: %v", err)
			}

			if !reflect.DeepEqual(r, game) {
				t.Errorf("was expecting\n%#v\nbut returns\n%#v\n", game, r)
			}
		})

		t.Run("failure", func(t *testing.T) {
			r, err := serviceFailure.Find(game.ID)
			if err == nil {
				t.Errorf(`was expecting an error, but returns: "%v"`, err)
			}

			if r != nil {
				t.Errorf("a result value in an error case should be nil, but returns: %v", r)
			}
		})
	})
}
